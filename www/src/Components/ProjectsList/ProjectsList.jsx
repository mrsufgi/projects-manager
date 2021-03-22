import * as React from "react";
import { DataGrid } from "@material-ui/data-grid";
import { Box } from "@material-ui/core";
import ReactJson from "react-json-view";
import useAxios from "axios-hooks";

const columns = [
  { field: "id", headerName: "ID", type: "number", width: 100 },
  { field: "name", headerName: "Name", width: 200 },
  { field: "vertical", headerName: "Vertical", width: 200 },
  { field: "event", headerName: "Event", width: 200 },
  { field: "url", headerName: "URL", width: 200 },
  {
    field: "credentials",
    headerName: "Credentials",
    width: 400,
    renderCell: (params) => {
      return (
        <Box style={{ maxHeight: "100%", overflow: "auto", width: "100%" }}>
          <ReactJson
            src={params.value}
            theme="bright:inverted"
            displayObjectSize={false}
            displayDataTypes={false}
            indentWidth={2}
          />
        </Box>
      );
    },
  },
];

const baseUrl = "http://localhost:4000"; // TODO: move to config

export default function DataTable() {
  const [{ data, loading, error }, refetch] = useAxios(`${baseUrl}/projects`);
  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error!</p>;
  return (
    <div style={{ width: "100%" }}>
      <DataGrid
        rows={data}
        columns={columns}
        pageSize={10}
        editMode="false"
        autoHeight
      />
    </div>
  );
}
