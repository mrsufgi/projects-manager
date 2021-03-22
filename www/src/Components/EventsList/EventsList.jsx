import * as React from "react";
import { DataGrid } from "@material-ui/data-grid";
import useAxios from "axios-hooks";

const columns = [
  { field: "id", headerName: "ID", type: "number", width: 100 },
  { field: "project_id", headerName: "Project ID", width: 150 },
  { field: "name", headerName: "Name", width: 350 },
  {
    field: "timestamp",
    headerName: "Timestamp",
    width: 350,
    type: "dateTime",
    sortable: true,
  },
];

const baseUrl = "http://localhost:4000"; // TODO: move to config

export default function ProjectsList() {
  const [{ data, loading, error }, refetch] = useAxios(`${baseUrl}/events`);
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
