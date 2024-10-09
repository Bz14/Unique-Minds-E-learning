"use client";
import { useState } from "react";

const RoleComponent = () => {
  const [userType, setUserType] = useState("student");
  return (
    <>
      <div className="flex space-x-4 justify-center">
        <button
          onClick={() => setUserType("student")}
          type="button"
          className={`w-full py-3 px-6 ${
            userType == "student"
              ? "bg-customBlue text-white"
              : "bg-white text-gray-800"
          } rounded-lg shadow-md hover:bg-gray-500 transition ease-in-out duration-300`}
        >
          Student
        </button>
        <button
          onClick={() => setUserType("teacher")}
          type="button"
          className={`w-full py-3 px-6 ${
            userType == "teacher"
              ? "bg-customBlue text-white"
              : "bg-white text-gray-800"
          } rounded-lg shadow-md hover:bg-gray-500 transition ease-in-out duration-300`}
        >
          Teacher
        </button>
      </div>
      ;
    </>
  );
};

export default RoleComponent;
