import React, { useState } from "react";
import StudentDashboard from "../Profile";
import Courses from "./Courses";
import {
  FaChartLine,
  FaBook,
  FaCalendarAlt,
  FaSignOutAlt,
  FaUser,
} from "react-icons/fa";
import ProgressSection from "../Progress";
import Schedule from "./Schedule";
import LogoutContent from "../../Educator/Logout";

// Function to display the sidebar of the student
const Sidebar = () => {
  const [activeSection, setActiveSection] = useState("Dashboard");

  const handleSectionClick = (section) => {
    setActiveSection(section);
  };

  return (
    <div className="flex w-full h-full">
      <div className="bg-customBlue text-white w-64 py-6 px-4 rounded-lg ml-3 mr-3 mt-1 h-full">
        <div className="flex flex-col items-center">
          <FaUser className="text-4xl mb-4" />
          <ul className="mt-4 space-y-4 w-full">
            <li
              className={`flex items-center space-x-2 cursor-pointer p-2 rounded-lg ${
                activeSection === "Dashboard" ? "bg-white text-customBlue" : ""
              }`}
              onClick={() => handleSectionClick("Dashboard")}
            >
              <FaUser
                className={activeSection === "Profile" ? "text-customBlue" : ""}
              />
              <span>Dashboard</span>
            </li>
            <li
              className={`flex items-center space-x-2 cursor-pointer p-2 rounded-lg ${
                activeSection === "Courses" ? "bg-white text-customBlue" : ""
              }`}
              onClick={() => handleSectionClick("Courses")}
            >
              <FaBook
                className={activeSection === "Courses" ? "text-customBlue" : ""}
              />
              <span>Courses</span>
            </li>
            <li
              className={`flex items-center space-x-2 cursor-pointer p-2 rounded-lg ${
                activeSection === "Progress" ? "bg-white text-customBlue" : ""
              }`}
              onClick={() => handleSectionClick("Progress")}
            >
              <FaChartLine
                className={
                  activeSection === "Progress" ? "text-customBlue" : ""
                }
              />
              <span>Progress</span>
            </li>
            <li
              className={`flex items-center space-x-2 cursor-pointer p-2 rounded-lg ${
                activeSection === "Schedule" ? "bg-white text-customBlue" : ""
              }`}
              onClick={() => handleSectionClick("Schedule")}
            >
              <FaCalendarAlt
                className={
                  activeSection === "Schedule" ? "text-customBlue" : ""
                }
              />
              <span>Schedule</span>
            </li>
            <li
              className={`flex items-center space-x-2 cursor-pointer p-2 rounded-lg ${
                activeSection === "Logout" ? "bg-white text-customBlue" : ""
              }`}
              onClick={() => handleSectionClick("Logout")}
            >
              <FaSignOutAlt
                className={activeSection === "Logout" ? "text-customBlue" : ""}
              />
              <span>Logout</span>
            </li>
          </ul>
        </div>
      </div>

      <div className="flex-1 flex flex-col mt-4 w-full">
        {activeSection === "Dashboard" && <StudentDashboard />}
        {activeSection === "Courses" && <Courses />}
        {activeSection === "Progress" && <ProgressSection />}
        {activeSection === "Schedule" && <Schedule />}
        {activeSection === "Logout" && <LogoutContent />}
      </div>
    </div>
  );
};

export default Sidebar;
