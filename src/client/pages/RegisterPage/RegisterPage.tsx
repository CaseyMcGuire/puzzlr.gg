import FormTextInput from "common/forms/FormTextInput";


export default function RegisterPage() {
  return (
    <div>
      <form action="/register" method="POST">
        <FormTextInput name={"email"} id={"email"} label={"Email"} />
        <FormTextInput id={"password"} name={"password"} label={"Password"} isPassword={true}/>
        <input type="submit"/>
      </form>
    </div>
  )
}