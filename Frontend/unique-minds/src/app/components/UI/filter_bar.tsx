import React, { useState } from "react";

const FilterBar = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [selectedFilters, setSelectedFilters] = useState({
    intelligence: false,
    language: false,
    speechTherapy: false,
    emotionalIntelligence: false,
  });

  interface Filters {
    intelligence: boolean;
    language: boolean;
    speechTherapy: boolean;
    emotionalIntelligence: boolean;
  }

  interface CheckboxChangeEvent {
    target: {
      name: string;
      checked: boolean;
    };
  }

  const handleCheckboxChange = (event: CheckboxChangeEvent) => {
    const { name, checked } = event.target;
    setSelectedFilters((prevState: Filters) => ({
      ...prevState,
      [name]: checked,
    }));
  };

  const handleApplyFilters = () => {
    console.log("Filters applied:", selectedFilters);
  };

  return (
    <div className="flex flex-col items-start shadow-2xl bg-white p-4">
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="block lg:hidden bg-customBlue text-white p-3 rounded-lg mb-4"
      >
        <svg
          className="h-6 w-6"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="2"
            d={isOpen ? "M6 18L18 6M6 6l12 12" : "M4 6h16M4 12h16M4 18h16"}
          ></path>
        </svg>
      </button>

      <div
        className={`flex flex-col space-x-4 mb-4 gap-6 ${
          isOpen ? "block" : "hidden"
        } lg:block`}
      >
        <h3>Filter By:</h3>
        <label className="flex items-center space-x-2 p-2">
          <input
            type="checkbox"
            name="intelligence"
            checked={selectedFilters.intelligence}
            onChange={handleCheckboxChange}
            className="appearance-none h-4 w-4 border-2 border-gray-400 checked:bg-customBlue checked:border-gray-400 focus:ring-1 focus:ring-customBlue"
          />
          <span>Intelligence</span>
        </label>
        <label className="flex items-center space-x-2 p-2">
          <input
            type="checkbox"
            name="language"
            checked={selectedFilters.language}
            onChange={handleCheckboxChange}
            className="appearance-none h-4 w-4 border-2 border-gray-400 checked:bg-customBlue checked:border-gray-400 focus:ring-1 focus:ring-customBlue"
          />
          <span>Language</span>
        </label>
        <label className="flex items-center space-x-2 p-2">
          <input
            type="checkbox"
            name="speechTherapy"
            checked={selectedFilters.speechTherapy}
            onChange={handleCheckboxChange}
            className="appearance-none h-4 w-4 border-2 border-gray-400 checked:bg-customBlue checked:border-gray-400 focus:ring-1 focus:ring-customBlue"
          />
          <span>Speech Therapy</span>
        </label>
        <label className="flex items-center space-x-2 p-2">
          <input
            type="checkbox"
            name="emotionalIntelligence"
            checked={selectedFilters.emotionalIntelligence}
            onChange={handleCheckboxChange}
            className="appearance-none h-4 w-4 border-2 border-gray-400 checked:bg-customBlue checked:border-gray-400 focus:ring-1 focus:ring-customBlue"
          />
          <span>Emotional Intelligence</span>
        </label>
        <button
          onClick={handleApplyFilters}
          className="bg-customBlue text-white p-3 hover:bg-gray-500 rounded-lg"
        >
          Apply Filter
        </button>
      </div>
    </div>
  );
};

export default FilterBar;
