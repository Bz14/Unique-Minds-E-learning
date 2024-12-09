import React, { useState, useEffect } from "react";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
  CartesianGrid,
  ResponsiveContainer,
} from "recharts";

const ProgressBarChart = () => {
  const sampleCourses = [];

  const [courses, setCourses] = useState(sampleCourses);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch("https://unique-minds.onrender.com/api/courses/progress", {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("access_token")}`,
      },
    })
      .then((response) => response.json())
      .then((data) => {
        setCourses(data);
        console.log(data);
        setLoading(false);
      })
      .catch((error) => {
        console.error("Error fetching course progress data:", error);
        setLoading(false);
      });
  }, []);

  return (
    <ResponsiveContainer width="100%" height={300}>
      <BarChart
        data={courses}
        margin={{ top: 20, right: 30, left: 20, bottom: 5 }}
      >
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="name" />
        <YAxis />
        <Tooltip />
        <Legend />
        <Bar dataKey="progress" fill="rgb(20, 74, 102)" />
      </BarChart>
    </ResponsiveContainer>
  );
};

// ProgressSection component
const ProgressSection = () => {
  return (
    <div className="w-full p-6">
      <h2 className="text-2xl font-bold mb-6 text-customBlue">
        Course Progress
      </h2>
      <ProgressBarChart />
    </div>
  );
};

export default ProgressSection;
