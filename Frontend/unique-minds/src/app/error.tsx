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
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <h1 className="text-3xl font-bold text-red-500">
        Oops! Something went wrong.
      </h1>
      <p className="mt-4 text-lg text-gray-700">{error.message}</p>

      <button
        onClick={reset} // This will reset the error boundary
        className="mt-6 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-500"
      >
        Try again
      </button>
    </div>
  );
};

export default Error;
