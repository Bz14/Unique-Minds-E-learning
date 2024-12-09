import React, { useState } from "react";
import {
  FaUser,
  FaBook,
  FaUpload,
  FaCalendarAlt,
  FaUsers,
  FaSignOutAlt,
} from "react-icons/fa";
import Profile from "../Components/Educator/Profile";
import Courses from "../Components/Educator/Courses";
import UploadCourse from "../Components/Educator/UploadCourses";
import Schedule from "../Components/Educator/Schedule";
import Students from "../Components/Educator/Students";
import LogoutContent from "../Components/Educator/Logout";
import Navbar from "../Components/Navbar";

// EducatorDashboard component
const EducatorDashboard = () => {
  const [activeSection, setActiveSection] = useState("profile");

  return (
    <>
      <Navbar />
      <div className="flex w-full h-full">
        <div className="bg-customBlue text-white w-64 py-6 px-4 h-screen rounded-lg ml-2 mr-2 mt-8">
          <div className="flex flex-col items-center">
            <FaUser className="text-4xl mb-4" />
            <ul className="mt-4 space-y-4 w-full">
              <li
                className={`flex items-center space-x-2 cursor-pointer p-2 rounded-lg ${
                  activeSection === "profile" ? "bg-white text-blue-600" : ""
                }`}
                onClick={() => setActiveSection("profile")}
              >
                <FaUser
                  className={activeSection === "profile" ? "text-blue-600" : ""}
                />
                <span>Profile</span>
              </li>
              <li
                className={`flex items-center space-x-2 cursor-pointer p-2 rounded-lg ${
                  activeSection === "courses" ? "bg-white text-blue-600" : ""
                }`}
                onClick={() => setActiveSection("courses")}
              >
                <FaBook
                  className={activeSection === "courses" ? "text-blue-600" : ""}
                />
                <span>Courses</span>
              </li>
              <li
                className={`flex items-center space-x-2 cursor-pointer p-2 rounded-lg ${
                  activeSection === "upload" ? "bg-white text-blue-600" : ""
                }`}
                onClick={() => setActiveSection("upload")}
              >
                <FaUpload
                  className={activeSection === "upload" ? "text-blue-600" : ""}
                />
                <span>Upload Course</span>
              </li>
              <li
                className={`flex items-center space-x-2 cursor-pointer p-2 rounded-lg ${
                  activeSection === "schedule" ? "bg-white text-blue-600" : ""
                }`}
                onClick={() => setActiveSection("schedule")}
              >
                <FaCalendarAlt
                  className={
                    activeSection === "schedule" ? "text-blue-600" : ""
                  }
                />
                <span>Schedule</span>
              </li>
              <li
                className={`flex items-center space-x-2 cursor-pointer p-2 rounded-lg ${
                  activeSection === "students" ? "bg-white text-blue-600" : ""
                }`}
                onClick={() => setActiveSection("students")}
              >
                <FaUsers
                  className={
                    activeSection === "students" ? "text-blue-600" : ""
                  }
                />
                <span>Students</span>
              </li>
              <li
                className={`flex items-center space-x-2 cursor-pointer p-2 rounded-lg ${
                  activeSection === "logout" ? "bg-white text-blue-600" : ""
                }`}
                onClick={() => setActiveSection("Logout")}
              >
                <FaSignOutAlt
                  className={activeSection === "Logout" ? "text-blue-600" : ""}
                />
                <span>Logout</span>
              </li>
            </ul>
          </div>
        </div>
        <div className="flex-1 flex flex-col p-8">
          {activeSection === "profile" && <Profile />}
          {activeSection === "courses" && <Courses />}
          {activeSection === "upload" && <UploadCourse />}
          {activeSection === "schedule" && <Schedule />}
          {activeSection === "students" && <Students />}
          {activeSection === "Logout" && <LogoutContent />}
        </div>
      </div>
      v
    </>
  );
};

export default EducatorDashboard;
