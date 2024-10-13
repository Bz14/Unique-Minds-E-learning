"use client";
import Link from "next/link";
import { FcGoogle } from "react-icons/fc";
import { MdVisibility, MdVisibilityOff } from "react-icons/md";
import { useForm } from "react-hook-form";
import { useEffect, useState } from "react";
import { DevTool } from "@hookform/devtools";
import { useRouter } from "next/navigation";

type SignUpForm = {
  fullName: string;
  email: string;
  password: string;
  confirmPassword: string;
  userType: string;
};

const SignUp = () => {
  const form = useForm<SignUpForm>({
    defaultValues: {
      fullName: "",
      email: "",
      password: "",
      confirmPassword: "",
      userType: "student",
    },
  });
  const { register, control, handleSubmit, setValue, formState, watch } = form;
  const { errors } = formState;
  const [passwordVisible, setPasswordVisible] = useState(true);
  const [password, setPassword] = useState("");
  const [confirmPasswordVisible, setConfirmPasswordVisible] = useState(true);
  const [userType, setUserType] = useState("student");
  const router = useRouter();

  useEffect(() => {
    const password = watch("password");
    setPassword(password);
  }, [watch("password")]);

  const onSubmit = async (data: SignUpForm) => {
    try {
      const response = await fetch("http://localhost:3000/api/auth/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      if (!response.ok) {
        const data = await response.json();
        throw new Error(
          data.message || "An error occurred while creating your account."
        );
      }
      router.push("/login");
    } catch (error) {
      console.log("Error", error);
      throw error;
    }
  };

  const handleUserType = (type: string) => {
    setUserType(type);
    setValue("userType", type);
  };
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="shadow-2xl rounded-lg overflow-hidden max-w-lg w-full bg-white mt-10 mx-auto">
        <div className="p-10 md:p-16">
          <h1 className="text-3xl font-bold mb-8 text-gray-800 text-center">
            Create Your Account
          </h1>
          <form
            className="space-y-6"
            onSubmit={handleSubmit(onSubmit)}
            noValidate
          >
            <div>
              <label
                htmlFor="full-name"
                className="block text-sm font-semibold text-gray-600"
              >
                Full Name
              </label>
              <input
                type="text"
                placeholder="Enter your full name"
                {...register("fullName", {
                  required: {
                    value: true,
                    message: "Full name is required.",
                  },
                })}
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
              <p style={{ color: "red", fontSize: "12px" }}>
                {errors.fullName?.message}
              </p>
            </div>
            <div>
              <label
                htmlFor="email"
                className="block text-sm font-semibold text-gray-600"
              >
                Email Address
              </label>
              <input
                {...register("email", {
                  required: {
                    value: true,
                    message: "Email is required.",
                  },
                  pattern: {
                    value:
                      /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/,
                    message: "Invalid email.",
                  },
                })}
                type="email"
                placeholder="Enter your email"
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
              <p style={{ color: "red", fontSize: "12px" }}>
                {errors.email?.message}
              </p>
            </div>
            <div className="relative">
              <label
                htmlFor="password"
                className="block text-sm font-semibold text-gray-600"
              >
                Password
              </label>
              <input
                {...register("password", {
                  required: {
                    value: true,
                    message: "Password is required.",
                  },
                  pattern: {
                    value:
                      /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/,
                    message:
                      "Password must be at least 8 characters long, include an uppercase letter, a lowercase letter, a number, and a special character",
                  },
                })}
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
              <p style={{ color: "red", fontSize: "12px" }}>
                {errors.password?.message}
              </p>
            </div>
            <div className="relative">
              <label
                htmlFor="confirm-password"
                className="block text-sm font-semibold text-gray-600"
              >
                Confirm Password
              </label>
              <input
                {...register("confirmPassword", {
                  required: {
                    value: true,
                    message: "Please confirm your password.",
                  },
                  validate: {
                    matchPassword: (value) => {
                      return value === password || "Passwords do not match.";
                    },
                  },
                })}
                type={confirmPasswordVisible ? "password" : "text"}
                placeholder="Confirm your password"
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
              <div
                className="absolute inset-y-0 right-3 flex items-center cursor-pointer"
                onClick={() =>
                  setConfirmPasswordVisible(!confirmPasswordVisible)
                }
              >
                <span className="mt-6">
                  {confirmPasswordVisible ? (
                    <MdVisibilityOff size={24} className="text-gray-500" />
                  ) : (
                    <MdVisibility size={24} className="text-gray-500" />
                  )}
                </span>
              </div>
              <p style={{ color: "red", fontSize: "12px" }}>
                {errors.confirmPassword?.message}
              </p>
            </div>
            <div className="flex space-x-4 justify-center">
              <button
                onClick={() => handleUserType("student")}
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
                onClick={() => handleUserType("teacher")}
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
            <div>
              <button
                type="submit"
                className="w-full py-3 px-6 bg-customBlue text-white rounded-lg shadow-md hover:bg-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition ease-in-out duration-300"
              >
                Create Account
              </button>
            </div>
            <div className="text-center mt-6">
              <h3 className="text-sm">
                Already have an account?{" "}
                <Link href="/user/login">
                  <span className="text-indigo-600 hover:text-indigo-500 font-semibold">
                    Log In
                  </span>
                </Link>
              </h3>
            </div>
            <div className="relative my-6">
              <div className="absolute inset-0 flex items-center">
                <div className="w-full border-t border-gray-300"></div>
              </div>
              <div className="relative flex justify-center text-sm">
                <span className="bg-white px-2 text-gray-500">Or</span>
              </div>
            </div>
            <div>
              <button
                type="button"
                className="w-full py-3 px-6 bg-white text-black rounded-lg shadow-md border border-gray-300 hover:bg-gray-200 transition ease-in-out duration-300"
              >
                <div className="flex align-center justify-center">
                  <span className="mt-1 mr-2">
                    <FcGoogle />
                  </span>
                  <span>Sign Up with Google</span>
                </div>
              </button>
            </div>
          </form>
        </div>
      </div>
      <DevTool control={control} />
    </div>
  );
};

export default SignUp;
