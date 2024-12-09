import React from "react";
import { Link } from "react-router-dom";

// CourseCard component
const CourseCard = ({ course }) => {
  return (
    <div className="border border-gray-200 rounded-lg shadow-lg overflow-hidden transition-transform transform hover:scale-105 max-w-sm mx-auto">
      <img
        src={course.image}
        alt={course.name}
        className="w-full h-40 object-cover"
      />
      <div className="p-4 bg-customBlue">
        <h3 className="text-2xl font-bold mb-2 text-white">{course.name}</h3>
        <p className="text-white mb-4">{course.description}</p>
        {console.log(`/courses/${course._id}`)}
        <Link
          to={`/courses/${course._id}`}
          className="inline-block bg-white text-customBlue py-2 px-4 rounded-lg hover:bg-gray-500 hover:text-white transition"
        >
          View Details
        </Link>
      </div>
    </div>
  );
};

export default CourseCard;
