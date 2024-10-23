"use client";
import { useState } from "react";
import Course from "../../Courses/course";
import img3 from "../../../assets/course.jpg";
import img1 from "../../../assets/communication.jpg";
import img2 from "../../../assets/intelligence.jpg";
import { StaticImageData } from "next/image";

const course_list = [
  {
    id: 1,
    title: "Speech and Communication Therapy",
    description:
      "This course include introduction to speech and communication therapy.",
    image: img1,
  },
  {
    id: 2,
    title: "Emotional Intelligence",
    description: "This course include introduction to emotional intelligence.",
    image: img2,
  },
  {
    id: 3,
    title: "Communication Skills",
    description: "This course include introduction to communication skills.",
    image: img3,
  },
];

interface Course {
  id: number;
  title: string;
  description: string;
  image: StaticImageData;
}

const FeaturedCourseSection = () => {
  const [courses, setCourses] = useState<Course[]>(course_list);

  return (
    <section className="featured-courses mt-10 px-4 md:px-16 lg:px-24">
      <h2 className="text-3xl font-bold text-center text-gray-800 mb-8">
        Featured Courses
      </h2>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        {courses.map((course) => (
          <Course course={course} key={course.id} />
        ))}
      </div>
    </section>
  );
};

export default FeaturedCourseSection;
