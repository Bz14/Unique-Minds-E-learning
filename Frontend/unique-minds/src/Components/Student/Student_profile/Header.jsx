import React from "react";
import { FaBell, FaSearch } from "react-icons/fa";

// Header component
const Header = () => {
  return (
    <header className="bg-white shadow p-4 flex justify-between items-center">
      <div className="flex items-center w-1/2 border rounded-md">
        <FaSearch className="text-gray-500 ml-2" />
        <input
          type="text"
          className="w-full px-4 py-2 focus:outline-none"
          placeholder="Search"
        />
      </div>

      <div className="flex items-center space-x-4">
        <div className="flex items-center space-x-2">
          <img
            src="https://randomuser.me/api/portraits/men/32.jpg"
            alt="User"
            className="w-10 h-10 rounded-full"
          />
          <div>
            <h4 className="text-gray-700 font-semibold">John Doe</h4>
            <p className="text-sm text-gray-500">3rd year</p>
          </div>
        </div>
        <button className="p-2 rounded-full bg-red-500 text-white">
          <FaBell />
        </button>
      </div>
    </header>
  );
};

export default Header;
