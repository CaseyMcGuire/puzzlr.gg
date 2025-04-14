import FormTextInput from "common/forms/FormTextInput";

export default function LoginPage() {
  return (
    <div>
      <form action="/login" method="POST">
        <FormTextInput name={"email"} id={"email"} label={"Email"} />
        <FormTextInput id={"password"} name={"password"} label={"Password"} />
        <input type="submit"/>
      </form>
    </div>
  )
}