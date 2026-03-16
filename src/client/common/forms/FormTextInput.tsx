
import * as stylex from "@stylexjs/stylex";

type FormTextInputProps = {
  id: string,
  name: string,
  label: string,
  isPassword?: boolean
}

const styles = stylex.create({
  field: {
    display: "flex",
    flexDirection: "column",
    gap: "6px",
  },
  label: {
    fontSize: "0.95rem",
    fontWeight: "500",
  },
  input: {
    width: "100%",
    borderWidth: "1px",
    borderStyle: "solid",
    borderColor: "#c9ced6",
    borderRadius: "6px",
    padding: "10px 12px",
    fontSize: "0.95rem",
    outline: "none",
    transitionProperty: "border-color, box-shadow",
    transitionDuration: "120ms",
    transitionTimingFunction: "ease-in-out",
    ":focus": {
      borderColor: "#2563eb",
      boxShadow: "0 0 0 3px rgba(37, 99, 235, 0.15)",
    },
  },
});

export default function FormTextInput(props: FormTextInputProps) {
  const type = (props.isPassword ?? false) ? "password" : "text";

  return (
    <div sx={styles.field}>
      <label sx={styles.label} htmlFor={props.id}>{props.label}</label>
      <input sx={styles.input} type={type} id={props.id} name={props.name} />
    </div>
  )
}
