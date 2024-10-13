"use client";
import { MdVisibility, MdVisibilityOff } from "react-icons/md";
import { useState } from "react";

const PasswordComponent = () => {
  const [passwordVisible, setPasswordVisible] = useState(true);
  const [confirmPasswordVisible, setConfirmPasswordVisible] = useState(true);
  return (
    <>
      <div className="relative">
        <label
          htmlFor="password"
          className="block text-sm font-semibold text-gray-600"
        >
          Password
        </label>
        <input
          name="password"
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
      </div>
      <div className="relative">
        <label
          htmlFor="confirm-password"
          className="block text-sm font-semibold text-gray-600"
        >
          Confirm Password
        </label>
        <input
          name="confirm-password"
          type={confirmPasswordVisible ? "password" : "text"}
          placeholder="Confirm your password"
          className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
        />
        <div
          className="absolute inset-y-0 right-3 flex items-center cursor-pointer"
          onClick={() => setConfirmPasswordVisible(!confirmPasswordVisible)}
        >
          <span className="mt-6">
            {confirmPasswordVisible ? (
              <MdVisibilityOff size={24} className="text-gray-500" />
            ) : (
              <MdVisibility size={24} className="text-gray-500" />
            )}
          </span>
        </div>
      </div>
    </>
  );
};

export default PasswordComponent;
