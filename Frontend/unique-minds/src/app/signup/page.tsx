import Link from "next/link";
import Image from "next/image";
import signup from "../assets/signup.jpg";
import { FcGoogle } from "react-icons/fc";

const SignUp = () => {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="shadow-2xl rounded-lg overflow-hidden grid grid-cols-1 md:grid-cols-2 max-w-6xl w-full bg-white mt-10">
        <div className="hidden md:block md:mt-24">
          <Image
            src={signup}
            alt="Signup Illustration"
            className="w-full h-full object-cover"
            layout="responsive"
          />
        </div>
        <div className="p-10 md:p-16">
          <h1 className="text-3xl font-bold mb-8 text-gray-800 text-center">
            Create Your Account
          </h1>
          <form className="space-y-6">
            <div>
              <label className="block text-sm font-semibold text-gray-600">
                Full Name
              </label>
              <input
                type="text"
                placeholder="Enter your full name"
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
            </div>

            <div>
              <label className="block text-sm font-semibold text-gray-600">
                Email Address
              </label>
              <input
                type="email"
                placeholder="Enter your email"
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
            </div>

            <div>
              <label className="block text-sm font-semibold text-gray-600">
                Password
              </label>
              <input
                type="password"
                placeholder="Create a password"
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
            </div>

            <div>
              <label className="block text-sm font-semibold text-gray-600">
                Confirm Password
              </label>
              <input
                type="password"
                placeholder="Confirm your password"
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
            </div>

            <div className="flex space-x-4 justify-center">
              <button
                type="button"
                className="w-full py-3 px-6 bg-customBlue text-white rounded-lg shadow-md hover:bg-blue-400 transition ease-in-out duration-300"
              >
                Student
              </button>
              <button
                type="button"
                className="w-full py-3 px-6 bg-white text-gray rounded-lg shadow-md hover:bg-gray-300 transition ease-in-out duration-300"
              >
                Teacher
              </button>
            </div>
            <div>
              <button
                type="submit"
                className="w-full py-3 px-6 bg-customBlue text-white rounded-lg shadow-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition ease-in-out duration-300"
              >
                Create Account
              </button>
            </div>

            {/* Already have an account */}
            <div className="text-center mt-6">
              <h3 className="text-sm">
                Already have an account?{" "}
                <Link href="/login">
                  <span className="text-indigo-600 hover:text-indigo-500 font-semibold">
                    Log In
                  </span>
                </Link>
              </h3>
            </div>

            {/* Divider */}
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
    </div>
  );
};

export default SignUp;
