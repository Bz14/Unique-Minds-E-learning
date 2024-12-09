import React, { useState } from "react";
import { FaEnvelope } from "react-icons/fa";
import { Link } from "react-router-dom";

// ForgotPassword component
const ForgotPassword = () => {
  const [email, setEmail] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await fetch(
      "https://unique-minds.onrender.com/api/auth/forgot-password",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email }),
      }
    );

    if (response.ok) {
      alert("A reset link has been sent to your email.");
    } else {
      alert("Failed to send reset link.");
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-md p-8 space-y-8 bg-white rounded shadow">
        <div>
          <h2 className="text-3xl font-bold text-center text-gray-900">
            Forgot Password
          </h2>
          <p className="mt-2 text-center text-gray-600">
            Enter your email to reset your password
          </p>
        </div>
        <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
          <div className="space-y-4">
            <div className="flex items-center border-b border-blue-500 py-2">
              <FaEnvelope className="text-blue-500 mr-3" />
              <input
                type="email"
                name="email"
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="w-full py-2 text-gray-700 leading-tight focus:outline-none"
              />
            </div>
          </div>

          <div>
            <button
              type="submit"
              className="w-full px-4 py-2 text-white bg-blue-500 rounded-md hover:bg-blue-600"
            >
              Send Reset Link
            </button>
          </div>

          <div className="text-sm text-center mt-4">
            Remembered your password?{" "}
            <Link to="/login" className="text-blue-500 hover:underline">
              Login
            </Link>
          </div>
        </form>
      </div>
    </div>
  );
};

export default ForgotPassword;
