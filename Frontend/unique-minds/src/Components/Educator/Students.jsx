import React, { useState, useEffect } from "react";
import { FaUser } from "react-icons/fa";

const sampleCoursesWithStudents = {};

// Function to display the students of the educator
const Students = () => {
  const [coursesWithStudents, setCoursesWithStudents] = useState(
    sampleCoursesWithStudents
  );
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchCoursesWithStudents = async () => {
      try {
        const token = localStorage.getItem("access_token");
        const response = await fetch(
          "https://unique-minds.onrender.com/api/educator/students",
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );

        if (!response.ok) {
          throw new Error("Failed to fetch courses and students");
        }

        const data = await response.json();
        console.log(data);
        setCoursesWithStudents(data);
      } catch (error) {
        setError(error.message);
      } finally {
        setLoading(false);
      }
    };

    fetchCoursesWithStudents();
  }, []);

  if (loading) {
    return <p className="text-white">Loading students data...</p>;
  }

  if (error) {
    return <p className="text-red-500">Error: {error}</p>;
  }

  return (
    <div className="bg-customBlue p-8 rounded-lg shadow-lg">
      <h2 className="text-3xl font-bold mb-6 text-white">Students Overview</h2>
      <div className="space-y-8">
        {coursesWithStudents != null &&
        Object.keys(coursesWithStudents).length > 0 ? (
          Object.keys(coursesWithStudents).map((courseName) => (
            <div key={courseName} className="bg-white p-6 rounded-lg shadow-md">
              <h3 className="text-2xl font-semibold text-customBlue mb-4">
                {courseName}
              </h3>
              <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
                {coursesWithStudents[courseName].length > 0 ? (
                  coursesWithStudents[courseName].map((student) => (
                    <div
                      key={student.id}
                      className="bg-customBlue text-white p-4 rounded-lg shadow-lg border border-gray-200 flex flex-col items-center space-y-4 transition-transform transform hover:scale-105"
                    >
                      <div className="w-16 h-16 bg-indigo-100 rounded-full flex items-center justify-center mb-4">
                        <FaUser className="text-3xl" />
                      </div>
                      <h4 className="text-lg font-semibold">{student.name}</h4>
                      <p className="text-sm">
                        {student.bio ? student.bio : "No bio available."}
                      </p>
                    </div>
                  ))
                ) : (
                  <p className="text-gray-500">No students for this course.</p>
                )}
              </div>
            </div>
          ))
        ) : (
          <p className="text-white">No courses or students found.</p>
        )}
      </div>
    </div>
  );
};

export default Students;
