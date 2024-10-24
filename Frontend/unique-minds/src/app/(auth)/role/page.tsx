"use client";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

const Role = () => {
  const apiUrl = process.env.NEXT_PUBLIC_API;
  const [role, setRole] = useState("student");
  const [email, setEmail] = useState("");
  const router = useRouter();

  useEffect(() => {
    const queryParam = new URLSearchParams(window.location.search);
    const email = queryParam.get("email");
    if (email) {
      setEmail(email);
    }
  }, []);

  const handleRole = async (role: string) => {
    try {
      const response = await fetch(
        `${apiUrl}/api/auth/role?email=${email}&&role=${role}`
      );
      if (!response.ok) {
        const data = await response.json();
        throw new Error(
          data.message || "An error occurred while logging you in."
        );
      }
      router.push(`/`);
    } catch (error) {
      throw error;
    }
  };

  return (
    <div className="flex flex-col items-center justify-center bg-white shadow-lg w-fit m-auto p-8 rounded-lg mt-20">
      <h1 className="text-3xl font-bold mb-8 text-gray-800">I am here as</h1>
      <div className="flex space-x-4">
        <button
          onClick={() => {
            setRole("student");
            handleRole("student");
          }}
          className={`px-6 py-3 rounded-lg text-white font-medium transition-all ${
            role === "student"
              ? "bg-customBlue hover:bg-blue-800"
              : "bg-gray-500 hover:bg-gray-400"
          }`}
        >
          Student
        </button>
        <button
          onClick={() => {
            setRole("teacher");
            handleRole("teacher");
          }}
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
