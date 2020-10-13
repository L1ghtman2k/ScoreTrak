package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/qor/validations"
	"gorm.io/gorm"
	"net/http"
	"reflect"
)

//Generic function passing and assignment

func genericGetByID(svc interface{}, log logger.LogInfoFormat, m string, idParam string, w http.ResponseWriter, r *http.Request) {
	id, err := uuidResolver(idParam, r)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sg, err := InvokeRetMethod(svc, m, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Error(err)
		}
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
	}
}

func genericGet(svc interface{}, log logger.LogInfoFormat, m string, w http.ResponseWriter, r *http.Request) {
	sg, err := InvokeRetMethod(svc, m)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Error(err)
		}
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
	}
}

func genericUpdate(svc interface{}, g interface{}, log logger.LogInfoFormat, m string, idParam string, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(g)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := uuidResolver(idParam, r)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	v := reflect.ValueOf(g).Elem()
	f := reflect.ValueOf(id)
	if _, ok := svc.(team.Serv); ok && idParam == "name" {
		v.FieldByName("Name").Set(f)
	} else {
		v.FieldByName("ID").Set(f)
	}
	err = InvokeNoRetMethod(svc, m, g)
	if err != nil {
		_, ok := err.(*validations.Error)
		if ok {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Error(err)
		return
	}
}

func genericStore(svc interface{}, g interface{}, log logger.LogInfoFormat, m string, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(g)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = InvokeNoRetMethod(svc, m, g)
	if err != nil {
		_, ok := err.(*validations.Error)
		if ok {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Error(err)
		return
	}
}

func genericDelete(svc interface{}, log logger.LogInfoFormat, m string, idParam string, w http.ResponseWriter, r *http.Request) {
	id, err := uuidResolver(idParam, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Error(err)
		return
	}
	err = InvokeNoRetMethod(svc, m, id)
	if err != nil {
		_, ok := err.(*orm.NoRowsAffected)
		if ok {
			http.Redirect(w, r, "/", http.StatusNotModified)
			return
		} else {
			http.Error(w, err.Error(), http.StatusConflict)
			log.Error(err)
			return
		}
	}
}

//Credit to:
// https://stackoverflow.com/questions/14116840/dynamically-call-method-on-interface-regardless-of-receiver-type
// https://stackoverflow.com/questions/8103617/call-a-struct-and-its-method-by-name-in-go
func InvokeRetMethod(i interface{}, methodName string, args ...interface{}) (interface{}, error) {
	finalMethod := PreInvoke(i, methodName)
	if finalMethod.IsValid() {
		inputs := make([]reflect.Value, len(args))
		for i := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		r := finalMethod.Call(inputs)

		if err, ok := r[1].Interface().(error); ok {
			return nil, err
		}
		return r[0].Interface(), nil
	}
	return nil, fmt.Errorf("the method name %s does not exist in %s", methodName, reflect.TypeOf(i).Name())
}

func InvokeNoRetMethod(i interface{}, methodName string, args ...interface{}) error {
	finalMethod := PreInvoke(i, methodName)
	if finalMethod.IsValid() {
		inputs := make([]reflect.Value, len(args))
		for i := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		r := finalMethod.Call(inputs)

		if err, ok := r[0].Interface().(error); ok {
			return err
		}
		return nil
	}
	return fmt.Errorf("the method name %s does not exist in %s", methodName, reflect.TypeOf(i).Name())
}

func PreInvoke(i interface{}, methodName string) reflect.Value {
	var ptr reflect.Value
	var value reflect.Value
	var finalMethod reflect.Value

	value = reflect.ValueOf(i)
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(i))
		temp := ptr.Elem()
		temp.Set(value)
	}
	method := value.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}
	method = ptr.MethodByName(methodName)
	if method.IsValid() {
		finalMethod = method
	}
	return finalMethod
}

func uuidResolver(idParam string, r *http.Request) (uuid.UUID, error) {
	params := mux.Vars(r)
	id := params[idParam]
	if id == "" {
		return uuid.Nil, errors.New("id was not found")
	}
	return uuid.FromString(id)
}
