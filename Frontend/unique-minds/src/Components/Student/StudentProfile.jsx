import Navbar from "../Navbar";
import SideBar from "./Student_profile/Sidebar";

// Student Profile Component
const StudentProfile = () => {
  return (
    <>
      <Navbar />
      <div className="flex h-screen bg-gray-100 w-full">
        <SideBar />
      </div>
    </>
  );
};

export default StudentProfile;
