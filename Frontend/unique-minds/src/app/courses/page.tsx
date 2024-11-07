"use client";
import SearchBar from "../components/UI/search_bar";
import FilterBar from "../components/UI/filter_bar";
import img from "../assets/img3.jpg";
import Course from "../components/Courses/course";
import { StaticImageData } from "next/image";

type Course = {
  id: number;
  image: StaticImageData;
  title: string;
  description: string;
};

const courses: Course[] = [
  {
    id: 1,
    image: img,
    title: "Intelligence",
    description: "This course is about intelligence.",
  },
  {
    id: 2,
    image: img,
    title: "Language",
    description: "This course is about language.",
  },
  {
    id: 3,
    image: img,
    title: "Speech Therapy",
    description: "This course is about speech therapy.",
  },
  {
    id: 4,
    image: img,
    title: "Emotional Intelligence",
    description: "This course is about emotional intelligence.",
  },
  {
    id: 1,
    image: img,
    title: "Intelligence",
    description: "This course is about intelligence.",
  },
  {
    id: 2,
    image: img,
    title: "Language",
    description: "This course is about language.",
  },
  {
    id: 3,
    image: img,
    title: "Speech Therapy",
    description: "This course is about speech therapy.",
  },
  {
    id: 4,
    image: img,
    title: "Emotional Intelligence",
    description: "This course is about emotional intelligence.",
  },
  {
    id: 1,
    image: img,
    title: "Intelligence",
    description: "This course is about intelligence.",
  },
  {
    id: 2,
    image: img,
    title: "Language",
    description: "This course is about language.",
  },
  {
    id: 3,
    image: img,
    title: "Speech Therapy",
    description: "This course is about speech therapy.",
  },
  {
    id: 4,
    image: img,
    title: "Emotional Intelligence",
    description: "This course is about emotional intelligence.",
  },
];

const Courses = () => {
  return (
    <div className="flex flex-col">
      <div className="mt-10">
        <SearchBar />
      </div>
      <div className="flex flex-row">
        <div className="w-1/4 mt-10 sticky top-0 h-screen bg-white p-4 shadow-lg">
          <FilterBar />
        </div>
        <div className="w-3/4 mt-10 p-4 overflow-y-auto">
          <div className="flex flex-row flex-wrap justify-center">
            {courses.map((course) => (
              <div key={course.id} className="w-72 mx-4 my-4">
                <Course course={course} />
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Courses;
