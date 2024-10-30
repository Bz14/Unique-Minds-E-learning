import React from "react";
import { MdSearch } from "react-icons/md";

const SearchBar = () => {
  return (
    <div className="flex items-center justify-center my-5">
      <div className="relative">
        <MdSearch className="w-5 h-5 text-customBlue absolute left-3 top-1/2 transform -translate-y-1/2" />
        <input
          type="text"
          placeholder="Search..."
          className="w-72 md:w-96 pl-10 pr-4 py-2.5 border-2 text-gray-500 border-gray-300 rounded-l-full focus:outline-none focus:border-customBlue transition-colors duration-300"
        />
      </div>
      <button className="px-6 py-2.5 bg-customBlue text-white font-semibold rounded-r-full hover:bg-gray-500 transition-colors duration-300">
        Search
      </button>
    </div>
  );
};

export default SearchBar;
