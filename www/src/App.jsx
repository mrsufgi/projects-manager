import "./App.css";

import React, { useState } from "react";
import Container from "@material-ui/core/Container";
import Typography from "@material-ui/core/Typography";
import Box from "@material-ui/core/Box";
import AddIcon from "@material-ui/icons/Add";

import EventsList from "./Components/EventsList/EventsList";
import ProjectsList from "./Components/ProjectsList/ProjectsList";
import AddProjectModal from "./Components/AddProjectModal/AddProjectModal";
import { Fab, makeStyles } from "@material-ui/core";

const useStyles = makeStyles((theme) => ({
  fab: {
    position: "absolute",
    bottom: theme.spacing(2),
    right: theme.spacing(2),
  },
}));

function App() {
  const classes = useStyles();
  const [open, setOpen] = useState(false);

  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <Container maxWidth="xl">
      <Fab
        color="primary"
        className={classes.fab}
        aria-label="add"
        onClick={handleOpen}
      >
        <AddIcon />
      </Fab>
      <Box my={12}>
        <Typography variant="h5" component="h1" gutterBottom>
          Event List
        </Typography>
        <EventsList></EventsList>
        <br />
        <br />
        <Typography variant="h5" component="h1" gutterBottom>
          Project List
        </Typography>
        <ProjectsList></ProjectsList>
        <AddProjectModal open={open} handleClose={handleClose} />
      </Box>
    </Container>
  );
}

export default App;
