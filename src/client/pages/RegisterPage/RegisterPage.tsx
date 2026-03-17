import FormTextInput from "common/forms/FormTextInput";
import * as stylex from "@stylexjs/stylex";

const styles = stylex.create({
  page: {
    minHeight: "100%",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    padding: "24px",
    backgroundColor: "#f6f7f9",
    color: "#1f2937",
    fontFamily: "Segoe UI, Tahoma, Geneva, Verdana, sans-serif",
  },
  cardForm: {
    width: "100%",
    maxWidth: "420px",
    borderWidth: "1px",
    borderStyle: "solid",
    borderColor: "#d1d5db",
    borderRadius: "8px",
    backgroundColor: "#ffffff",
    padding: "24px",
    display: "flex",
    flexDirection: "column",
    gap: "14px",
  },
  title: {
    marginBottom: "16px",
    fontSize: "1.5rem",
    fontWeight: "600",
  },
  submit: {
    marginTop: "4px",
    border: "none",
    borderRadius: "6px",
    padding: "10px 14px",
    backgroundColor: "#2563eb",
    color: "#ffffff",
    fontSize: "0.95rem",
    fontWeight: "600",
    cursor: "pointer",
    transitionProperty: "background-color",
    transitionDuration: "120ms",
    transitionTimingFunction: "ease-in-out",
    ":hover": {
      backgroundColor: "#1d4ed8",
    },
  },
});


export default function RegisterPage() {
  return (
    <div sx={styles.page}>
      <form sx={styles.cardForm} action="/user/create" method="POST">
        <h1 sx={styles.title}>Register</h1>
        <FormTextInput name={"email"} id={"email"} label={"Email"} />
        <FormTextInput id={"password"} name={"password"} label={"Password"} isPassword={true}/>
        <input sx={styles.submit} type="submit" value="Register" />
      </form>
    </div>
  )
}
