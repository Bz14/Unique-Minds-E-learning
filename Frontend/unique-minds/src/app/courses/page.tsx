import SearchBar from "../components/UI/search_bar";
import FilterBar from "../components/UI/filter_bar";

const Courses = () => {
  return (
    <div className="flex flex-col">
      <div className="mt-10">
        <SearchBar />
      </div>
      <div className="mt-10">
        <FilterBar />
      </div>
      All Courses
    </div>
  );
};

export default Courses;
