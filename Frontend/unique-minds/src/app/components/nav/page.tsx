"use client";
import Link from "next/link";
import { useState } from "react";
import Image from "next/image";

const Nav = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

  return (
    <nav className="bg-customBlue px-4 py-2 shadow-md">
      <div className="container mx-auto flex items-center justify-between">
        <div className="text-white font-bold text-xl">
          <Link href="/" className="flex items-center">
            <Image
              src="/logo.png"
              alt="Logo"
              width={32}
              height={32}
              className="mr-2"
            />
            <span>YourLogo</span>
          </Link>
        </div>

        <div className="block lg:hidden">
          <button
            className="text-white focus:outline-none"
            onClick={toggleMenu}
          >
            <svg
              className="h-6 w-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d={isOpen ? "M6 18L18 6M6 6l12 12" : "M4 6h16M4 12h16M4 18h16"}
              ></path>
            </svg>
          </button>
        </div>

        <div
          className={`lg:flex items-center lg:space-x-6 ${
            isOpen ? "block" : "hidden"
          } lg:block`}
        >
          <ul className="flex flex-col lg:flex-row lg:space-x-6 space-y-2 lg:space-y-0">
            <li>
              <Link href="/courses" className="text-white hover:text-gray-300">
                Courses
              </Link>
            </li>
            <li>
              <Link href="/teacher" className="text-white hover:text-gray-300">
                Teacher
              </Link>
            </li>
            <li>
              <Link
                href="/dashboard"
                className="text-white hover:text-gray-300"
              >
                Dashboard
              </Link>
            </li>
            <li>
              <Link href="/about" className="text-white hover:text-gray-300">
                About Us
              </Link>
            </li>
            <li>
              <Link href="/login" className="text-white hover:text-gray-300">
                Login
              </Link>
            </li>
            <li>
              <Link
                href="/signup"
                className="bg-white text-customBlue px-4 py-2 rounded-md hover:bg-gray-100"
              >
                Sign Up
              </Link>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
};

export default Nav;
