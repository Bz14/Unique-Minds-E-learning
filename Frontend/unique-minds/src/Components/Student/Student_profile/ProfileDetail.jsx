import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

// Function to log out the student
const StudentDashboard = () => {
  const [isEditing, setIsEditing] = useState(false);
  const [profile, setProfile] = useState({
    profileImage: "https://via.placeholder.com/150",
    name: "John Doe",
    age: "16",
    condition: "Down Syndrome",
    bio: "Student bio goes here...",
    guardianEmail: "guardian@example.com",
    guardianPhone: "123-456-7890",
    location: "City, Country",
  });

  const navigate = useNavigate();

  useEffect(() => {
    const fetchProfile = async () => {
      try {
        const response = await fetch(
          "https://unique-minds.onrender.com/api/profile",
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
            },
          }
        );
        if (!response.ok) {
          throw new Error("Failed to fetch profile");
        }
        const data = await response.json();
        setProfile(data);
      } catch (error) {
        console.error("Error fetching profile data:", error);
      }
    };

    fetchProfile();
  }, []);

  const handleEditClick = () => {
    if (isEditing) {
      fetch("https://unique-minds.onrender.com/api/studentProfile", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        body: JSON.stringify(profile),
      })
        .then((response) => response.json())
        .then((data) => {
          setProfile(data);
          setIsEditing(false);
        })
        .catch((error) => console.error("Error posting profile data:", error));
    } else {
      setIsEditing(true);
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setProfile((prevProfile) => ({
      ...prevProfile,
      [name]: value,
    }));
  };

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        setProfile((prevProfile) => ({
          ...prevProfile,
          profileImage: reader.result,
        }));
      };
      reader.readAsDataURL(file);
    }
  };

  return (
    <main className="p-6 bg-gray-100 flex-1">
      <div className="bg-customBlue p-6 rounded-lg mb-6">
        <h1 className="text-2xl font-bold text-white">
          Welcome back, {profile.name}!
        </h1>
        <p className="mt-2 text-white">
          Always stay updated in your student portal
        </p>
      </div>

      <div className="bg-light-gray p-8 rounded shadow-md w-full">
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-3xl font-bold text-black">Student Profile</h2>
          <button
            onClick={handleEditClick}
            className="text-blue-600 hover:text-blue-800"
          >
            {isEditing ? "Save" : "Edit"}
          </button>
        </div>

        <div className="flex items-center space-x-6">
          <label htmlFor="imageUpload" className="cursor-pointer">
            <div className="w-32 h-32">
              <img
                src={profile.profileImage}
                alt="Profile"
                className="w-full h-full rounded-full object-cover shadow-lg"
              />
            </div>
            {isEditing && (
              <input
                id="imageUpload"
                type="file"
                className="hidden"
                onChange={handleImageChange}
              />
            )}
          </label>

          <div className="flex-grow">
            {isEditing ? (
              <input
                type="text"
                name="name"
                value={profile.name}
                onChange={handleChange}
                className="text-3xl font-bold text-black bg-white border border-gray-300 rounded p-2 mb-2"
              />
            ) : (
              <h2 className="text-3xl font-bold text-black">{profile.name}</h2>
            )}
            <div className="flex items-center space-x-4">
              {isEditing ? (
                <input
                  type="text"
                  name="age"
                  value={profile.age}
                  onChange={handleChange}
                  className="text-gray-600 bg-white border border-gray-300 rounded p-2"
                />
              ) : (
                <p className="text-gray-600">Age: {profile.age}</p>
              )}
              {isEditing ? (
                <input
                  type="text"
                  name="condition"
                  value={profile.condition}
                  onChange={handleChange}
                  className="text-gray-600 bg-white border border-gray-300 rounded p-2"
                />
              ) : (
                <p className="text-gray-600">Condition: {profile.condition}</p>
              )}
            </div>
          </div>
        </div>

        <div className="mt-6">
          <h3 className="text-2xl font-bold text-black mb-4">Bio</h3>
          {isEditing ? (
            <textarea
              name="bio"
              value={profile.bio}
              onChange={handleChange}
              className="w-full h-32 bg-white border border-gray-300 rounded p-2 text-gray-800 mb-4"
            />
          ) : (
            <p className="text-gray-800">{profile.bio}</p>
          )}
        </div>

        <div className="mt-6">
          <h3 className="text-2xl font-bold text-black mb-4">
            Guardian Details
          </h3>
          <div className="mb-4">
            <h4 className="text-lg font-semibold text-black mb-2">Email</h4>
            {isEditing ? (
              <input
                type="email"
                name="guardianEmail"
                value={profile.guardianEmail}
                onChange={handleChange}
                className="w-full bg-white border border-gray-300 rounded p-2"
              />
            ) : (
              <p>{profile.guardianEmail}</p>
            )}
          </div>

          <div className="mb-4">
            <h4 className="text-lg font-semibold text-black mb-2">Phone</h4>
            {isEditing ? (
              <input
                type="text"
                name="guardianPhone"
                value={profile.guardianPhone}
                onChange={handleChange}
                className="w-full bg-white border border-gray-300 rounded p-2"
              />
            ) : (
              <p>{profile.guardianPhone}</p>
            )}
          </div>

          <div className="mb-4">
            <h4 className="text-lg font-semibold text-black mb-2">Location</h4>
            {isEditing ? (
              <input
                type="text"
                name="location"
                value={profile.location}
                onChange={handleChange}
                className="w-full bg-white border border-gray-300 rounded p-2"
              />
            ) : (
              <p>{profile.location}</p>
            )}
          </div>
        </div>
      </div>
    </main>
  );
};

export default StudentDashboard;
