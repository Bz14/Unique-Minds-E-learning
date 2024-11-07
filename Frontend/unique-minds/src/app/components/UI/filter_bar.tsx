import React, { useState } from "react";

const FilterBar = () => {
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
    // Logic for applying the filters (e.g., fetch data or update UI)
    console.log("Filters applied:", selectedFilters);
  };

  return (
    <div className="flex flex-col items-start shadow-2xl bg-white p-4">
      <div className="flex flex-col space-x-4 mb-4 gap-6">
        <h3>Filter By:</h3>
        <label className="flex items-center space-x-2">
          <input
            type="checkbox"
            name="intelligence"
            checked={selectedFilters.intelligence}
            onChange={handleCheckboxChange}
            className="appearance-none h-4 w-4 border-2 border-gray-400  checked:bg-customBlue checked:border-gray-400 focus:ring-1 focus:ring-customBlue"
          />
          <span>Intelligence</span>
        </label>
        <label className="flex items-center space-x-2">
          <input
            type="checkbox"
            name="language"
            checked={selectedFilters.language}
            onChange={handleCheckboxChange}
            className="appearance-none h-4 w-4 border-2 border-gray-400  checked:bg-customBlue checked:border-gray-400 focus:ring-1 focus:ring-customBlue"
          />
          <span>Language</span>
        </label>
        <label className="flex items-center space-x-2">
          <input
            type="checkbox"
            name="speechTherapy"
            checked={selectedFilters.speechTherapy}
            onChange={handleCheckboxChange}
            className="appearance-none h-4 w-4 border-2 border-gray-400  checked:bg-customBlue checked:border-gray-400 focus:ring-1 focus:ring-customBlue"
          />
          <span>Speech Therapy</span>
        </label>
        <label className="flex items-center space-x-2">
          <input
            type="checkbox"
            name="emotionalIntelligence"
            checked={selectedFilters.emotionalIntelligence}
            onChange={handleCheckboxChange}
            className="appearance-none h-4 w-4 border-2 border-gray-400  checked:bg-customBlue checked:border-gray-400 focus:ring-1 focus:ring-customBlue"
          />
          <span>Emotional Intelligence</span>
        </label>
      </div>
      <button
        onClick={handleApplyFilters}
        className="bg-customBlue text-white p-3 hover:bg-gray-500 rounded-lg ml-auto"
      >
        Apply Filter
      </button>
    </div>
  );
};

export default FilterBar;
