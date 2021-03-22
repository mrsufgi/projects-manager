import * as React from "react";
import { Formik, Form, Field } from "formik";
import * as Yup from "yup";
import { Button, LinearProgress } from "@material-ui/core";
import { TextField } from "formik-material-ui";
import Box from "@material-ui/core/Box";
import useAxios from "axios-hooks";

const addSchema = Yup.object().shape({
  name: Yup.string()
    .min(2, "Too Short!")
    .max(50, "Too Long!")
    .required("Required"),
  vertical: Yup.string()
    .min(2, "Too Short!")
    .max(50, "Too Long!")
    .required("Required"),
  event: Yup.string()
    .min(2, "Too Short!")
    .max(100, "Too Long!")
    .required("Required"),
  url: Yup.string()
    .min(2, "Too Short!")
    .max(100, "Too Long!")
    .required("Required"),
  credentials: Yup.object(),
});

const baseUrl = "http://localhost:4000";
export default function AddProjectForm(props) {
  const [{ loading, error }, createProject] = useAxios(
    {
      method: "POST",
      baseURL: baseUrl,
      url: "/projects",
    },
    { manual: true }
  );

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error!</p>;
  const { handleClose } = props;
  return (
    <Formik
      validationSchema={addSchema}
      initialValues={{
        name: "",
        vertical: "",
        event: "",
        url: "",
        credentials: undefined,
      }}
      onSubmit={async (values, { setSubmitting }) => {
        const { credentials, ...rest } = values;
        const parsedCredentials = credentials
          ? JSON.parse(credentials)
          : undefined;
        await createProject({
          data: { ...rest, credentials: parsedCredentials },
        });
        setSubmitting(false);
        handleClose();
      }}
    >
      {({ submitForm, isSubmitting, touched, errors }) => (
        <Form>
          <Box margin={1}>
            <Field component={TextField} name="name" type="text" label="Name" />
          </Box>
          <Box margin={1}>
            <Field
              component={TextField}
              name="vertical"
              type="text"
              label="Vertical"
            />
          </Box>
          <Box margin={1}>
            <Field
              component={TextField}
              name="event"
              type="text"
              label="Event"
            />
          </Box>
          <Box margin={1}>
            <Field component={TextField} name="url" type="text" label="URL" />
          </Box>
          <Box margin={1}>
            <Field
              component={TextField}
              name="credentials"
              type="text"
              label="Credentials"
              helperText="Enter JSON string"
            />
          </Box>

          {isSubmitting && <LinearProgress />}
          <Box margin={1}>
            <Button
              variant="contained"
              color="primary"
              disabled={isSubmitting}
              onClick={submitForm}
            >
              Submit
            </Button>
          </Box>
        </Form>
      )}
    </Formik>
  );
}
