import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import Navbar from "../Components/Navbar";

// CourseDetail component
const CourseDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [course, setCourse] = useState({
    id: "course_12345",
    name: "Introduction to Web Development",
    description:
      "This course will teach you the basics of web development, including HTML, CSS, and JavaScript.",
    image: "https://example.com/course-image.jpg",
    parts: [
      {
        id: "part_1",
        name: "Getting Started with HTML",
        description:
          "Learn the basics of HTML and how to structure your web pages.",
        materials: [
          {
            name: "HTML Cheat Sheet",
            type: "PDF",
            content: "https://example.com/materials/html-cheat-sheet.pdf",
          },
          {
            name: "HTML Tutorial Video",
            type: "Video",
            content: "https://example.com/materials/html-tutorial.mp4",
          },
        ],
      },
      {
        id: "part_2",
        name: "Introduction to CSS",
        description: "Learn how to style your web pages with CSS.",
        materials: [
          {
            name: "CSS Flexbox Guide",
            type: "PDF",
            content: "https://example.com/materials/css-flexbox-guide.pdf",
          },
          {
            name: "CSS Tutorial Video",
            type: "Video",
            content: "https://example.com/materials/css-tutorial.mp4",
          },
        ],
      },
      {
        id: "part_3",
        name: "Getting Started with JavaScript",
        description:
          "Learn the basics of JavaScript and how to make your web pages interactive.",
        materials: [
          {
            name: "JavaScript Basics",
            type: "PDF",
            content: "https://example.com/materials/js-basics.pdf",
          },
          {
            name: "JavaScript Tutorial Video",
            type: "Video",
            content: "https://example.com/materials/js-tutorial.mp4",
          },
        ],
      },
    ],
  });
  const [progressData, setProgress] = useState({
    id: "progress_12345",
    progress: 60,
    completed_parts: ["part_1"],
    is_completed: false,
  });

  const handleSaveCourse = async () => {
    const token = localStorage.getItem("access_token");

    if (!token) {
      //navigate("/login");
      return;
    }

    try {
      await fetch(`https://unique-minds.onrender.com/api/courses/${id}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      alert("Course saved successfully!");
    } catch (error) {
      console.error("Error saving course:", error);
    }
  };

  useEffect(() => {
    const fetchCourse = async () => {
      const token = localStorage.getItem("access_token");

      if (!token) {
        //navigate("/login");
        return;
      }

      try {
        const response = await fetch(
          `https://unique-minds.onrender.com/api/courses/${id}`,
          {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
              Authorization: `Bearer ${token}`,
            },
          }
        );

        const data = await response.json();

        if (response.ok) {
          setCourse(data.course);
          setProgress(data.progress);
        } else {
          console.error("Failed to fetch course");
        }
      } catch (error) {
        console.error("Error fetching course:", error);
      }
    };

    fetchCourse();
  }, [id, navigate]);

  const togglePartCompletion = async (partId) => {
    const updatedCompletedParts =
      progressData.completed_parts != null &&
      progressData.completed_parts.includes(partId)
        ? progressData.completed_parts.filter((id) => id !== partId)
        : progressData.completed_parts != null
        ? [...progressData.completed_parts, partId]
        : [];

    setProgress({ ...progressData, completed_parts: updatedCompletedParts });

    try {
      const token = localStorage.getItem("access_token");
      const response = await fetch(
        `https://unique-minds.onrender.com/api/courses/progress/${id}`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
          body: JSON.stringify(updatedCompletedParts),
        }
      );
      const data = await response.json();

      if (response.ok) {
        setProgress(data);
      } else {
        console.error("Failed to fetch course");
      }
    } catch (error) {
      console.error("Error updating progress:", error);
    }
  };

  return (
    <>
      <Navbar />
      <div className="container mx-auto my-12 px-4 mb-20 w-3/4">
        <div className="bg-white rounded-lg shadow-lg overflow-hidden">
          <img
            src={course.image}
            alt={course.name}
            className="w-full h-64 object-cover rounded-t-lg"
          />

          <div className="p-6">
            <h1 className="text-4xl font-extrabold text-customBlue mb-4">
              {course.name}
            </h1>

            <p className="text-lg text-gray-700 mb-8 leading-relaxed">
              {course.description}
            </p>

            <h2 className="text-2xl font-semibold text-customBlue mb-4 border-b-2 border-gray-200 pb-2">
              Course Parts
            </h2>

            <div className="space-y-6">
              {course.parts !== undefined &&
                course.parts.map((part) => (
                  <div
                    key={part._id}
                    className={`p-6 rounded-lg shadow-md ${
                      progressData.completed_parts !== null &&
                      progressData.completed_parts.includes(part._id)
                        ? "bg-customBlue text-white"
                        : "bg-gray-100 text-customBlue"
                    }`}
                  >
                    <div className="flex items-center">
                      <input
                        type="checkbox"
                        checked={
                          progressData.completed_parts !== null &&
                          progressData.completed_parts.includes(part._id)
                        }
                        onChange={() => togglePartCompletion(part._id)}
                        className="mr-3"
                      />
                      <h3 className="text-xl font-semibold mb-2">
                        {part.name}
                      </h3>
                    </div>
                    <p className="mb-4">{part.description}</p>

                    <div className="mt-4">
                      <h4 className="text-lg font-semibold mb-2">Materials</h4>
                      <ul className="list-disc list-inside pl-5">
                        {part.materials.map((material, index) => (
                          <li key={index}>
                            <a
                              href={material.content}
                              className="hover:underline"
                              target="_blank"
                              rel="noopener noreferrer"
                            >
                              {material.name} ({material.type})
                            </a>
                          </li>
                        ))}
                      </ul>
                    </div>
                  </div>
                ))}
            </div>

            <div className="flex justify-between mt-8">
              <div className="text-lg font-semibold">
                Progress: {progressData.progress}%
              </div>

              <button
                className="bg-gray-200 text-customBlue font-bold py-2 px-6 rounded-lg shadow hover:bg-gray-400 transition duration-300"
                onClick={handleSaveCourse}
              >
                Save Course
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default CourseDetail;
