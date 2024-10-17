"use client";
import { useState } from "react";
import { MdVisibility, MdVisibilityOff } from "react-icons/md";
// import { DevTool } from "@hookform/devtools";

const Login = () => {
  const [passwordVisible, setPasswordVisible] = useState(false);
  return (
    <div className="min-h-screen flex items-center justify-center bg-white">
      <div className="shadow-2xl rounded-lg overflow-hidden max-w-lg w-full bg-white mx-auto">
        <div className="p-10 md:p-16">
          <h1 className="text-3xl font-bold mb-8 text-gray-800 text-center">
            Welcome Back!
          </h1>
          <form className="space-y-6">
            <div>
              <label
                htmlFor="email"
                className="block text-sm font-semibold text-gray-600"
              >
                Email Address
              </label>
              <input
                type="email"
                placeholder="Enter your email"
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
              {/* <p style={{ color: "red", fontSize: "12px" }}>
                {errors.email?.message}
              </p> */}
            </div>
            <div className="relative">
              <label
                htmlFor="password"
                className="block text-sm font-semibold text-gray-600"
              >
                Password
              </label>
              <input
                type={passwordVisible ? "password" : "text"}
                placeholder="Create a password"
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
              <div
                className="absolute inset-y-0 right-3 flex items-center cursor-pointer"
                onClick={() => setPasswordVisible(!passwordVisible)}
              >
                <span className="mt-6">
                  {passwordVisible ? (
                    <MdVisibilityOff size={24} className="text-gray-500" />
                  ) : (
                    <MdVisibility size={24} className="text-gray-500" />
                  )}
                </span>
              </div>
              {/* <p style={{ color: "red", fontSize: "12px" }}>
                {errors.password?.message}
              </p> */}
            </div>

            <div className="flex items-center justify-between">
              <div className="flex items-center">
                <input
                  type="checkbox"
                  id="remember"
                  className="rounded border-gray-300 text-customBlue shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                />
                <label
                  htmlFor="remember"
                  className="ml-2 block text-sm text-gray-900"
                >
                  Remember me
                </label>
              </div>
            </div>
            <div>
              <button
                type="submit"
                className="w-full bg-customBlue text-white p-3 rounded-lg font-semibold hover:bg-gray-400"
              >
                Login
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Login;

{
  /* <DevTool control={control} /> */
}
