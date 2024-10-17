// create a simple page showing that a verification email has been sent to the user email and
// a button says check email when clicked it opens an email
import { MdEmail } from "react-icons/md";

const Verify = () => {
  return (
    <div className="flex align-middle justify-center mt-10 mx-auto shadow-lg w-fit max-h-screen bg-white">
      <div className="flex flex-col justify-center align-middle p-10 text-gray-800">
        <div className="bg-white shadow-lg mx-auto rounded-full">
          <MdEmail className="text-5xl mx-auto p-2" />
        </div>
        <h1 className="text-3xl font-bold text-gray-800 text-center p-5">
          Please verify your email
        </h1>
        <div className="flex flex-col justify-center align-middle mx-auto">
          <h4 className="text-xl text-center">
            We have sent a verification link to bezueyerusalem@gmail.com
          </h4>
          <p className="text-center p-4">
            Click on the link to complete the verification process. If you
            <br />
            don&apos;t see it, you may need to <b>check your spam folder.</b>
          </p>
          <h5 className="text-center font-bold p-2">
            Still can&apos;t find the email? No problem.
          </h5>
          <button className="bg-customBlue text-white w-fit mx-auto p-3 rounded-lg mt-2">
            Resend Email
          </button>
        </div>
      </div>
    </div>
  );
};

export default Verify;
