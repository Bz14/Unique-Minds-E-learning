import React, { useEffect, useState } from "react";
import CourseCard from "./CourseCard";

const Courses = () => {
  const [courses, setCourses] = useState();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const response = await fetch(
          "https://unique-minds.onrender.com/api/courses/my",
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem("access_token")}`,
            },
          }
        );
        if (!response.ok) {
          throw new Error("Failed to fetch courses");
        }
        const data = await response.json();
        console.log(data);
        setCourses(data);
        setLoading(false);
      } catch (err) {
        setError(err.message);
        setLoading(false);
      }
    };

    fetchCourses();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div className="w-full p-6 bg-white">
      <h2 className="text-2xl font-bold mb-6 text-customBlue">Your Courses</h2>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        {courses !== undefined &&
          courses.map((course) => (
            <CourseCard
              key={course._id}
              title={course.name}
              description={course.description}
              imageUrl={course.image}
              courseId={course._id}
            />
          ))}
      </div>
    </div>
  );
};

export default Courses;
