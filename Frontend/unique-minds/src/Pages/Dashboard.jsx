import React, { useState } from "react";
import Profile from "../Components/Student/Profile";
import Courses from "../Components/Student/Courses";
import Schedule from "../Components/Student/Schedule";
import Progress from "../Components/Student/Progress";
import Navbar from "../Components/Navbar";

// StudentProfile component
function StudentProfile() {
  const [currentTab, setCurrentTab] = useState(0);

  const handleTabChange = (index) => {
    setCurrentTab(index);
  };

  return (
    <>
      <Navbar />
      <div className="h-screen w-screen bg-gray-200 flex items-center justify-center">
        <div className="h-full w-full bg-white shadow-lg rounded-lg flex flex-col">
          <div className="flex bg-blue-600 text-white">
            <button
              className={`py-4 flex-1 text-center focus:outline-none ${
                currentTab === 0 ? "bg-gray-300" : "bg-blue-600"
              }`}
              onClick={() => handleTabChange(0)}
            >
              Profile
            </button>
            <button
              className={`py-4 flex-1 text-center focus:outline-none ${
                currentTab === 1 ? "bg-gray-300" : "bg-blue-600"
              }`}
              onClick={() => handleTabChange(1)}
            >
              Courses
            </button>
            <button
              className={`py-4 flex-1 text-center focus:outline-none ${
                currentTab === 2 ? "bg-gray-300" : "bg-blue-600"
              }`}
              onClick={() => handleTabChange(2)}
            >
              Schedule
            </button>
            <button
              className={`py-4 flex-1 text-center focus:outline-none ${
                currentTab === 3 ? "bg-gray-300" : "bg-blue-600"
              }`}
              onClick={() => handleTabChange(3)}
            >
              Progress
            </button>
          </div>
          <div className="p-6 flex-1 overflow-auto">
            {currentTab === 0 && <Profile />}
            {currentTab === 1 && <Courses />}
            {currentTab === 2 && <Schedule />}
            {currentTab === 3 && <Progress />}
          </div>
        </div>
      </div>
    </>
  );
}

export default StudentProfile;
