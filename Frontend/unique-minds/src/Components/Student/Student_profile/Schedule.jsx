import React, { useState, useEffect } from "react";
import { FaCalendarAlt, FaUser, FaGoogle } from "react-icons/fa";
import { MdCancel } from "react-icons/md";

// Function to display the schedule of the student
const Schedule = () => {
  const [schedules, setSchedules] = useState([]);

  useEffect(() => {
    const fetchSchedules = async () => {
      const token = localStorage.getItem("access_token");
      try {
        const response = await fetch(
          "https://unique-minds.onrender.com/api/student/schedules",
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );
        if (!response.ok) {
          throw new Error("Failed to fetch schedules");
        }
        const data = await response.json();
        setSchedules(data);
      } catch (error) {
        console.error("Error fetching schedules:", error);
      }
    };

    fetchSchedules();
  }, []);

  const handleCancelSchedule = async (id) => {
    try {
      const token = localStorage.getItem("access_token");
      const response = await fetch(
        `https://unique-minds.onrender.com/api/students/schedules/${id}`,
        {
          method: "DELETE",
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      if (!response.ok) {
        throw new Error("Failed to cancel schedule");
      }
      setSchedules(schedules.filter((schedule) => schedule.id !== id));
    } catch (error) {
      console.error("Error canceling schedule:", error);
    }
  };

  return (
    <div className="bg-white p-8 rounded-lg shadow-lg w-full h-full">
      <h2 className="text-4xl font-bold mb-6 text-customBlue">My Schedule</h2>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8">
        {schedules != null && schedules.length > 0 ? (
          schedules.map((schedule) => (
            <div
              key={schedule.id}
              className="flex flex-col justify-between bg-gray-50 p-6 rounded-lg shadow hover:shadow-md transition-shadow duration-300 min-h-full"
            >
              <div className="flex items-center space-x-4">
                <div className="bg-customBlue p-3 rounded-full">
                  <FaUser className="text-white text-2xl" />
                </div>
                <div className="flex flex-col space-y-1">
                  <h3 className="text-xl font-bold text-gray-800">
                    {schedule.educatorName}
                  </h3>
                  <div className="flex items-center space-x-2 text-gray-500">
                    <FaCalendarAlt />
                    <p>{schedule.date}</p>
                  </div>
                  <div className="flex items-center space-x-2 text-gray-500">
                    <FaGoogle />
                    <u>
                      <a
                        target="_blank"
                        href={schedule.googleMeetLink}
                        rel="noreferrer"
                        className="underline hover:text-customBlue"
                      >
                        Google Meet
                      </a>
                    </u>
                  </div>
                </div>
              </div>
              <button
                onClick={() => handleCancelSchedule(schedule.id)}
                className="mt-4 py-2 px-4 bg-red-500 text-white rounded-lg hover:bg-red-600 transition duration-300"
              >
                <MdCancel className="inline-block mr-1" />
                Cancel
              </button>
            </div>
          ))
        ) : (
          <p className="text-gray-700">No schedules available.</p>
        )}
      </div>
    </div>
  );
};

export default Schedule;
