import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { FaSearch } from "react-icons/fa";
import Navbar from "../Components/Navbar";

//
const Educators = () => {
  const [educators, setEducators] = useState([]);

  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [searchQuery, setSearchQuery] = useState("");

  const fetchEducators = async (query = "", page = 1) => {
    try {
      const response = await fetch(
        `https://unique-minds.onrender.com/api/educators?search=${query}&page=${page}&limit=6`
      );
      const data = await response.json();
      if (response.ok) {
        setEducators((prevEducators) =>
          page === 1
            ? data.data.educators
            : [...prevEducators, ...data.data.educators]
        );
        setTotalPages(data.data.pagination.totalPages);
      } else {
        console.error("Failed to fetch educators");
      }
    } catch (error) {
      console.error("Error fetching educators:", error);
    }
  };
  useEffect(() => {
    fetchEducators();
  }, []);

  const handleSearch = (query) => {
    setSearchQuery(query);
    fetchEducators(query, 1);
  };

  const loadMoreEducators = () => {
    if (page < totalPages) {
      const nextPage = page + 1;
      setPage(nextPage);
      fetchEducators(searchQuery, nextPage);
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
              placeholder="Search Educator"
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

  return (
    <>
      <Navbar />
      <SearchBar />
      <div className="container mx-auto my-8 px-4 mb-16">
        <h2 className="text-4xl font-bold text-center text-customBlue mt-12 mb-10">
          Educators
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 w-3/4 mx-auto">
          {educators != null &&
            educators.length > 0 &&
            educators.map((educator) => (
              <div
                key={educator._id}
                className="bg-customBlue rounded-lg shadow-lg overflow-hidden transition-transform transform hover:scale-105 hover:shadow-xl"
              >
                <div className="p-6 text-center">
                  <img
                    src={educator.profileImage}
                    alt={educator.name}
                    className="w-24 h-24 rounded-full mx-auto mb-4 object-cover"
                  />
                  <h3 className="text-2xl font-bold text-white mb-2">
                    {educator.name}
                  </h3>
                  <h4 className="text-lg text-white mb-2">{educator.title}</h4>
                  <button className="mt-4 bg-white text-customBlue font-semibold py-2 px-4 rounded-lg hover:bg-gray-500 focus:outline-none">
                    <Link to={`/educator_detail/${educator.id}`}>
                      View Profile
                    </Link>
                  </button>
                </div>
              </div>
            ))}
        </div>

        {page < totalPages && (
          <div className="flex justify-center mt-8">
            <button
              onClick={loadMoreEducators}
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

export default Educators;
