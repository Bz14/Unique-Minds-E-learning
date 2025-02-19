import React from "react";
import { useNavigate } from "react-router-dom";

const VerifyEmail = () => {
  const navigate = useNavigate();

  const handleGoToLogin = () => {
    navigate("/login");
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-md p-8 space-y-8 bg-white rounded shadow">
        <div>
          <h2 className="text-3xl font-bold text-center text-customBlue">
            Verify Your Email
          </h2>
          <p className="mt-2 text-center text-gray-600">
            A verification link has been sent to your email. Please check your
            inbox and follow the instructions to verify your account.
          </p>
        </div>
        <div className="mt-6 text-center">
          <button
            onClick={handleGoToLogin}
            className="px-4 py-2 text-white bg-customBlue rounded-md hover:bg-gray-500"
          >
            Go to Login
          </button>
        </div>
      </div>
    </div>
  );
};

export default VerifyEmail;
