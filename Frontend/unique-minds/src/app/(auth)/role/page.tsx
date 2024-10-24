"use client";
import { useState } from "react";

const Role = () => {
  const [role, setRole] = useState("student");

  return (
    <div className="flex flex-col items-center justify-center bg-white shadow-lg w-fit m-auto p-8 rounded-lg mt-20">
      <h1 className="text-3xl font-bold mb-8 text-gray-800">I am here as</h1>
      <div className="flex space-x-4">
        <button
          onClick={() => setRole("student")}
          className={`px-6 py-3 rounded-lg text-white font-medium transition-all ${
            role === "student"
              ? "bg-customBlue hover:bg-blue-800"
              : "bg-gray-500 hover:bg-gray-400"
          }`}
        >
          Student
        </button>
        <button
          onClick={() => setRole("teacher")}
          className={`px-6 py-3 rounded-lg text-white font-medium transition-all ${
            role === "teacher"
              ? "bg-customBlue hover:bg-blue-800"
              : "bg-gray-500 hover:bg-gray-400"
          }`}
        >
          Teacher
        </button>
      </div>
    </div>
  );
};

export default Role;
