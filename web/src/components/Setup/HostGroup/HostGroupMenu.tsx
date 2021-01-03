import React, {useEffect} from "react";
import {SetupProps} from "../util/util";
import Box from "@material-ui/core/Box";
import MaterialTable from "material-table";
import {Severity} from "../../../types/types";
import {CircularProgress} from "@material-ui/core";
import {UUID} from "../../../grpc/pkg/proto/utilpb/uuid_pb";
import {
    DeleteRequest,
    GetAllRequest,
    HostGroup,
    StoreRequest,
    UpdateRequest
} from "../../../grpc/pkg/host_group/host_grouppb/host_group_pb";
import {BoolValue} from "google-protobuf/google/protobuf/wrappers_pb";

type hostGroupColumns = {
    id: string | undefined
    enabled: boolean | undefined
    name: string,
}

function hostGroupColumnsToHostGroup(hostGroupC: hostGroupColumns): HostGroup{
    const t = new HostGroup()
    if (hostGroupC.enabled !== undefined) t.setEnabled(new BoolValue().setValue(hostGroupC.enabled))
    if (hostGroupC.id && hostGroupC.id !== "") t.setId((new UUID().setValue(hostGroupC.id)))
    t.setName(hostGroupC.name)
    return t
}

function hostGroupToHostGroupColumns(hostGroup: HostGroup): hostGroupColumns{
    return {
        enabled: hostGroup.getEnabled()?.getValue(),
        id: hostGroup.getId()?.getValue(),
        name: hostGroup.getName()
    }
}



export default function HostGroupsMenu(props: SetupProps) {
    const title = "Host Groups"
    props.setTitle(title)
    const columns=
        [
            { title: 'ID (optional)', field: 'id', editable: 'onAdd'},
            { title: 'Host Group Name', field: 'name' },
            { title: 'Enabled', field: 'enabled', type: 'boolean', initialEditValue: true},
        ]

    const [state, setState] = React.useState<{columns: any[], loader: boolean, data: hostGroupColumns[]}>({
        columns: columns,
        loader: true,
        data: []
    });

    function reloadSetter() {
        props.gRPCClients.hostGroupClient.getAll(new GetAllRequest(), {}).then(hostGroupsResponse => {
            setState(prevState => {return{...prevState, data: hostGroupsResponse.getHostGroupsList().map((hostGroup):hostGroupColumns => {
                    return hostGroupToHostGroupColumns(hostGroup)}), loader:false}})}, (err: any) => {
            props.genericEnqueue(`Encountered an error while retrieving Host Groups: ${err.message}. Error code: ${err.code}`, Severity.Error)
        })
    }
    useEffect(() => {
        reloadSetter()
    }, []);

    return (
        <React.Fragment>
            {!state.loader ?
                <Box height="100%" width="100%" >
                    <MaterialTable

                        editable={{
                            onRowAdd: (newData) =>
                                new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        const storeRequest = new StoreRequest()
                                        const u = hostGroupColumnsToHostGroup(newData)
                                        storeRequest.addHostGroups(u, 0)
                                        props.gRPCClients.hostGroupClient.store(storeRequest, {}).then(result => {
                                            u.setId(result.getIdsList()[0])
                                            setState((prevState) => {
                                                const data = [...prevState.data];
                                                data.push(hostGroupToHostGroupColumns(u));
                                                return { ...prevState, data };
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Unable to store hostGroup: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                            reject()
                                        })
                                    }, 600);
                                }),
                            onRowUpdate: (newData, oldData) =>
                                new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        if (oldData){
                                            const updateRequest = new UpdateRequest()
                                            const u = hostGroupColumnsToHostGroup(newData)
                                            updateRequest.setHostGroup(u)
                                            props.gRPCClients.hostGroupClient.update(updateRequest, {}).then(result => {
                                                setState((prevState) => {
                                                    const data = [...prevState.data];
                                                    data[data.indexOf(oldData)] = newData;
                                                    return { ...prevState, data };
                                                });
                                                resolve();
                                            }, (err: any) => {
                                                props.genericEnqueue(`Unable to update hostGroup: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                                reject()
                                            })
                                        }
                                    }, 600);
                                }),
                            onRowDelete: (oldData) =>
                                new Promise((resolve, reject) => {
                                    setTimeout(() => {
                                        const deleteRequest = new DeleteRequest()
                                        deleteRequest.setId((new UUID().setValue(oldData.id as string)))
                                        props.gRPCClients.hostGroupClient.delete(deleteRequest, {}).then(result => {
                                            setState((prevState) => {
                                                const data = [...prevState.data];
                                                data.splice(data.indexOf(oldData), 1);
                                                return { ...prevState, data };
                                            });
                                            resolve();
                                        }, (err: any) => {
                                            props.genericEnqueue(`Unable to delete hostGroup: ${err.message}. Error code: ${err.code}`, Severity.Error)
                                            reject()
                                        })
                                    }, 600);
                                }),
                        }}
                        
                        title={title}
                        columns={state.columns}
                        data={state.data}
                        options={{pageSizeOptions: [5,10,20,50,100, 500, 1000], pageSize:20, emptyRowsWhenPaging:false}}
                    />
                </Box>
                :
                <Box height="100%" width="100%" m="auto">
                    <CircularProgress  />
                </Box>
            }
        </React.Fragment>
    );
}