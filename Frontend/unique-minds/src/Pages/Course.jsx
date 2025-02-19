import React, { useEffect, useState } from "react";
import CourseCard from "../Components/Home/CourseCard";
import { FaSearch } from "react-icons/fa";
import Navbar from "../Components/Navbar";

// Courses component
const Courses = () => {
  const [courses, setCourses] = useState([]);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [searchQuery, setSearchQuery] = useState("");
  const [selectedTag, setSelectedTag] = useState("");

  const fetchCourses = async (query = "", tag = "", page = 1) => {
    try {
      const response = await fetch(
        `https://unique-minds.onrender.com/api/courses?search=${query}&tag=${tag}&page=${page}&limit=10`
      );
      const data = await response.json();

      if (response.ok) {
        setCourses((prevCourses) =>
          page === 1 ? data.data.course : [...prevCourses, ...data.data.course]
        );
        setTotalPages(data.data.pagination.totalPages);
      } else {
        console.error("Failed to fetch courses");
      }
    } catch (error) {
      console.error("Error fetching courses:", error);
    }
  };

  useEffect(() => {
    fetchCourses();
  }, []);

  const handleSearch = (query) => {
    setSearchQuery(query);
    fetchCourses(query, selectedTag, 1);
  };

  const handleTagSelect = (tag) => {
    setSelectedTag(tag);
    fetchCourses(searchQuery, tag, 1);
  };

  const loadMoreCourses = () => {
    if (page < totalPages) {
      const nextPage = page + 1;
      setPage(nextPage);
      fetchCourses(searchQuery, selectedTag, nextPage);
    }
  };

  const SearchBar = () => {
    const [input, setInput] = useState("");

    const handleInputChange = (e) => {
      setInput(e.target.value);
    };

    const handleSearchSubmit = (e) => {
      e.preventDefault();
      handleSearch(input);
    };

    return (
      <div className="flex justify-center pt-8 pb-4 px-4 mt-10 mb-20">
        <form
          onSubmit={handleSearchSubmit}
          className="w-full max-w-2xl flex items-center bg-white rounded-md overflow-hidden"
        >
          <div className="relative flex-1">
            <input
              type="text"
              value={input}
              onChange={handleInputChange}
              placeholder="Enter Course Name"
              className="w-full py-3 px-4 pl-12 rounded-md text-customBlue border border-blue-300 focus:outline-none"
            />
            <FaSearch className="absolute left-4 top-1/2 transform -translate-y-1/2 text-customBlue" />
          </div>
          <button
            type="submit"
            className="bg-customBlue text-white py-3 px-6 rounded-md hover:bg-gray-500 focus:outline-none"
          >
            Search
          </button>
        </form>
      </div>
    );
  };

  const FilterTags = () => {
    const tags = [
      "Speech",
      "Intelligence",
      "Physical Exercise",
      "Mental growth",
    ];

    return (
      <div className="flex flex-wrap justify-around gap-6 p-4 bg-white rounded-lg shadow-md mt-10">
        {tags.length > 0 &&
          tags.map((tag) => (
            <button
              key={tag}
              onClick={() => handleTagSelect(tag)}
              className="px-6 py-2 rounded-full bg-customBlue text-white border border-blue-300 hover:bg-gray-500 focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-300 ease-in-out"
            >
              {tag}
            </button>
          ))}
      </div>
    );
  };

  return (
    <>
      <Navbar />
      <SearchBar />
      <FilterTags />
      <div className="container mx-auto my-8 px-4 mb-16">
        <h2 className="text-4xl font-bold text-center text-customBlue mt-12 mb-10">
          Courses
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {courses != null &&
            courses.map((course) => (
              <div
                key={course._id}
                className="bg-white rounded-lg shadow-lg overflow-hidden transition-transform transform hover:scale-105 hover:shadow-xl"
              >
                <CourseCard course={course} key={course._id} />
              </div>
            ))}
        </div>
        {page < totalPages && (
          <div className="flex justify-center mt-8">
            <button
              onClick={loadMoreCourses}
              className="bg-blue-500 text-white py-2 px-4 rounded-lg hover:bg-blue-600 focus:outline-none"
            >
              Load More
            </button>
          </div>
        )}
      </div>
    </>
  );
};

export default Courses;
