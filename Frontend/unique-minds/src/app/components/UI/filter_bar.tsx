import React from "react";

const FilterBar = () => {
  return (
    <div className="flex flex-row align-middle justify-around shadow-2xl bg-white p-2">
      <button className="bg-customBlue text-white p-3 hover:bg-gray-500 rounded-lg">
        Intelligence
      </button>
      <button className="bg-customBlue text-white p-3 hover:bg-gray-500 rounded-lg">
        Language
      </button>
      <button className="bg-customBlue text-white p-3 hover:bg-gray-500 rounded-lg">
        Speech Therapy
      </button>
      <button className="bg-customBlue text-white p-3 hover:bg-gray-500 rounded-lg">
        Emotional Intelligence
      </button>
    </div>
  );
};

export default FilterBar;
