import React, { useState, useEffect, useRef } from "react";
import { FaEnvelope, FaPhoneAlt, FaMapMarkerAlt } from "react-icons/fa";

// Function to display the student dashboard
const StudentDashboard = () => {
  const [isEditing, setIsEditing] = useState(false);
  const [profile, setProfile] = useState({});
  const [imageFile, setImageFile] = useState(null);

  const fileInputRef = useRef(null);

  useEffect(() => {
    const fetchProfile = async () => {
      try {
        const response = await fetch(
          "https://unique-minds.onrender.com/api/profile",
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem("access_token")}`,
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

  useEffect(() => {
    if (imageFile) {
      const handleImageUpload = async () => {
        const formData = new FormData();
        formData.append("file", imageFile);

        try {
          const response = await fetch(
            "https://unique-minds.onrender.com/api/profile/upload",
            {
              method: "POST",
              body: formData,
              headers: {
                Accept: "application/json",
                Authorization: `Bearer ${localStorage.getItem("access_token")}`,
              },
            }
          );

          if (!response.ok) {
            throw new Error("Error uploading image");
          }

          const data = await response.json();
          setProfile((prevProfile) => ({
            ...prevProfile,
            profileImage: data.fileUrl,
          }));
          setImageFile(null);
        } catch (error) {
          console.error("Error uploading image:", error);
        }
      };

      handleImageUpload();
    }
  }, [imageFile]);

  const handleEditClick = () => {
    if (isEditing) {
      fetch("https://unique-minds.onrender.com/api/profile", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("access_token")}`,
        },
        body: JSON.stringify(profile),
      })
        .then((response) => response.json())
        .then((data) => {
          setProfile(data);
          setIsEditing(false);
        })
        .catch((error) => console.error("Error updating profile data:", error));
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

  const handleProfileImageClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleFileChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      setImageFile(file);
    }
  };

  return (
    <main className="p-8 bg-white flex-1">
      <div className="bg-customBlue p-8 rounded-xl mb-8 shadow-lg">
        <h1 className="text-3xl font-semibold text-white">
          Welcome back, {profile.name}!
        </h1>
        <p className="mt-3 text-white">
          Stay updated and track your progress on your student dashboard
        </p>
      </div>

      <div className="bg-white p-10 rounded-lg shadow-md w-full">
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-3xl font-bold text-customBlue">
            Student Profile
          </h2>
          <button
            onClick={handleEditClick}
            className="text-blue-500 hover:text-blue-700 font-medium"
          >
            {isEditing ? "Save" : "Edit"}
          </button>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label htmlFor="imageUpload" className="cursor-pointer">
              <div className="w-36 h-36 mb-6">
                <img
                  src={profile.profileImage}
                  alt="Profile"
                  className="w-full h-full rounded-full object-cover shadow-lg"
                  onClick={handleProfileImageClick}
                  style={{ cursor: "pointer" }}
                />
              </div>
              {isEditing && (
                <input
                  ref={fileInputRef}
                  id="imageUpload"
                  type="file"
                  accept="image/*"
                  className="hidden"
                  onChange={handleFileChange}
                />
              )}
            </label>

            <div className="flex-grow">
              {isEditing ? (
                <input
                  type="text"
                  name="name"
                  value={profile.name || ""}
                  onChange={handleChange}
                  className="text-3xl font-semibold text-customBlue bg-white border border-gray-300 rounded-lg p-3 mb-3 w-full"
                />
              ) : (
                <h2 className="text-3xl font-semibold text-customBlue">
                  {profile.name}
                </h2>
              )}

              <div className="flex space-x-4 mb-6">
                {isEditing ? (
                  <input
                    type="text"
                    name="age"
                    value={profile.age || ""}
                    onChange={handleChange}
                    className="text-customBlue bg-white border border-gray-300 rounded-lg p-2"
                  />
                ) : (
                  <p className="text-customBlue">Age: {profile.age}</p>
                )}

                {isEditing ? (
                  <input
                    type="text"
                    name="condition"
                    value={profile.condition || ""}
                    onChange={handleChange}
                    className="text-customBlue bg-white border border-gray-300 rounded-lg p-2"
                  />
                ) : (
                  <p className="text-customBlue">
                    Condition: {profile.condition}
                  </p>
                )}
              </div>

              <div className="mt-8">
                <h3 className="text-2xl font-bold text-customBlue mb-3">Bio</h3>
                {isEditing ? (
                  <textarea
                    name="bio"
                    value={profile.bio || ""}
                    onChange={handleChange}
                    className="w-full h-32 bg-white border border-gray-300 rounded-lg p-3 text-gray-800"
                  />
                ) : (
                  <p className="text-customBlue">{profile.bio}</p>
                )}
              </div>
            </div>
          </div>

          <div className="bg-customBlue rounded-lg shadow-lg w-fit h-fit p-10 text-white ml-10">
            <h3 className="text-2xl font-bold text-white mb-3">
              Guardian Details
            </h3>

            <div className="mb-4">
              {isEditing ? (
                <input
                  type="email"
                  name="guardianEmail"
                  value={profile.guardianEmail || ""}
                  onChange={handleChange}
                  className="w-full bg-white border border-gray-300 rounded-lg p-3 text-customBlue"
                  placeholder="Email"
                />
              ) : (
                <span className="flex items-center">
                  <FaEnvelope className=" mr-2" />
                  {profile.guardianEmail}
                </span>
              )}
            </div>

            <div className="mb-4">
              {isEditing ? (
                <input
                  type="text"
                  name="guardianPhone"
                  value={profile.guardianPhone || ""}
                  onChange={handleChange}
                  placeholder="Phone"
                  className="w-full bg-white border border-gray-300 rounded-lg p-3 text-customBlue"
                />
              ) : (
                <span className="flex items-center">
                  <FaPhoneAlt className=" mr-2" />
                  {profile.guardianPhone}
                </span>
              )}
            </div>

            <div className="mb-4">
              {isEditing ? (
                <input
                  type="text"
                  name="location"
                  value={profile.location || ""}
                  onChange={handleChange}
                  placeholder="Address"
                  className="w-full bg-white border border-gray-300 rounded-lg p-3 text-customBlue"
                />
              ) : (
                <span className="flex items-center">
                  <FaMapMarkerAlt className=" mr-2" />
                  {profile.location}
                </span>
              )}
            </div>
          </div>
        </div>
      </div>
    </main>
  );
};

export default StudentDashboard;
