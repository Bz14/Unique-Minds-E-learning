import { DevTool } from "@hookform/devtools";

const Login = () => {
  return (
    <div>
      <h1>Login</h1>
      <DevTool control={control} />
    </div>
  );
};

export default Login;
