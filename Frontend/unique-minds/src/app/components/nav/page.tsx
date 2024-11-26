"use client";
// import { Link, useLocation } from "react-router-dom";
// import { useEffect, useState } from "react";
import { useState } from "react";
import logo from "../../assets/logo.png";
import Image from "next/image";
import Link from "next/link";
import defaultImg from "../../assets/image0_0.jpg";

const Navbar = () => {
  const [accessToken, setAccessToken] = useState();
  const [profileImage, setProfileImage] = useState(defaultImg);
  const [role, setRole] = useState();
  // const location = useLocation(); // Get current location

  // useEffect(() => {
  //   const token = localStorage.getItem("access_token");

  //   if (token) {
  //     setAccessToken(token);
  //     fetchUserProfile(token);
  //   }
  // }, []);

  // useEffect(() => {
  //   // If the current route is the homepage, reload the page
  //   console.log(location.pathname);
  //   if (location.pathname === "/") {
  //     const hasReloaded = sessionStorage.getItem("hasReloadedHome");

  //     if (!hasReloaded) {
  //       sessionStorage.setItem("hasReloadedHome", "true");
  //       window.location.reload();
  //     }
  //   }
  // }, [location]); // Run the effect whenever the route changes

  // const fetchUserProfile = async (token) => {
  //   try {
  //     const response = await fetch("/api/auth/user-profile", {
  //       headers: {
  //         Authorization: `Bearer ${token}`,
  //       },
  //     });

  //     if (response.ok) {
  //       const data = await response.json();
  //       const profileImageUrl = data.profileImage
  //         ? data.profileImage
  //         : defaultImg;
  //       setProfileImage(profileImageUrl);
  //       setRole(data.role);
  //     }
  //   } catch (error) {
  //     console.error("Error fetching profile:", error);
  //   }
  // };

  return (
    <nav className="bg-white shadow-md p-4">
      <div className="container mx-auto flex justify-between items-center">
        <Link href="/" className="text-customBlue text-lg font-bold">
          <Image src={logo} alt="UniqueMinds" className="w-20 h-10" />
        </Link>
        <div className="flex items-center">
          <Link href="/courses" className="text-customBlue mr-4">
            Courses
          </Link>
          <Link href="/educators" className="text-customBlue mr-4">
            Educators
          </Link>
          <Link href="/footer" className="text-customBlue mr-4">
            Contact Us
          </Link>

          {!accessToken ? (
            <>
              <Link href="/login" className="text-customBlue mr-4">
                Login
              </Link>
              <Link
                href="/signup"
                className="bg-customBlue text-white px-4 py-2 rounded mr-4"
              >
                Signup
              </Link>
            </>
          ) : (
            <>
              {role === "student" && (
                <Link href="/student_dashboard" className="mr-4">
                  <div className="flex items-center">
                    <Image
                      src={profileImage}
                      alt="Profile"
                      className="w-10 h-10 rounded-full object-cover"
                    />
                  </div>
                </Link>
              )}
              {role === "educator" && (
                <Link
                  href="/educator_dashboard"
                  className="text-customBlue mr-4"
                >
                  <div className="flex items-center">
                    <Image
                      src={profileImage}
                      alt="Profile"
                      className="w-10 h-10 rounded-full object-cover"
                    />
                  </div>
                </Link>
              )}
            </>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
