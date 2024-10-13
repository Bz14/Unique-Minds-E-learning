"use client";

import { useEffect } from "react";

interface ErrorProps {
  error: Error;
  reset: () => void;
}

const Error = ({ error, reset }: ErrorProps) => {
  useEffect(() => {
    console.error(error);
  }, [error]);

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-50 p-4">
      <div className="bg-white shadow-lg rounded-lg p-6 max-w-md w-full text-center">
        <h1 className="text-3xl font-extrabold text-red-600 mb-2">
          Oops! Something went wrong.
        </h1>
        <p className="mt-2 text-lg text-gray-700">
          {error.message || "An unexpected error occurred."}
        </p>

        <button
          onClick={reset}
          className="mt-6 px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-500 transition duration-300 transform hover:scale-105"
        >
          Try Again
        </button>
      </div>
    </div>
  );
};

export default Error;
