import React, { useState } from "react";

// Function to upload a new course
const UploadCourse = () => {
  const [course, setCourse] = useState({
    name: "",
    description: "",
    image: null,
    isFeatured: false,
    parts: [],
    tags: [],
  });

  const [currentPart, setCurrentPart] = useState({
    name: "",
    description: "",
    sequence: 1,
    materials: [],
  });

  const [currentMaterial, setCurrentMaterial] = useState({
    name: "",
    type: "",
    content: null,
    description: "",
  });

  const [selectedPartSequence, setSelectedPartSequence] = useState(null);
  const [newTag, setNewTag] = useState("");

  const handleCourseChange = (e) => {
    const { name, value, type, checked, files } = e.target;
    setCourse((prevCourse) => ({
      ...prevCourse,
      [name]: type === "checkbox" ? checked : files ? files[0] : value,
    }));
  };

  const handlePartChange = (e) => {
    const { name, value } = e.target;
    setCurrentPart((prevPart) => ({
      ...prevPart,
      [name]: value,
    }));
  };

  const handleMaterialChange = (e) => {
    const { name, value, files } = e.target;
    setCurrentMaterial((prevMaterial) => ({
      ...prevMaterial,
      [name]: files ? files[0] : value,
    }));
  };

  const addPart = () => {
    setCourse((prevCourse) => ({
      ...prevCourse,
      parts: [...prevCourse.parts, currentPart],
    }));
    setCurrentPart({
      name: "",
      description: "",
      sequence: course.parts.length + 1,
      materials: [],
    });
  };

  const addMaterial = () => {
    if (selectedPartSequence === null) {
      alert("Please select a part to add materials.");
      return;
    }

    setCourse((prevCourse) => {
      const updatedParts = prevCourse.parts.map((part) =>
        part.sequence === selectedPartSequence
          ? { ...part, materials: [...part.materials, currentMaterial] }
          : part
      );
      return { ...prevCourse, parts: updatedParts };
    });

    setCurrentMaterial({
      name: "",
      type: "",
      content: null,
      description: "",
    });
  };

  const uploadFile = async (file) => {
    console.log("Uploading file:", file);
    const formData = new FormData();
    formData.append("file", file);

    try {
      const response = await fetch(
        "https://unique-minds.onrender.com/api/upload",
        {
          method: "POST",
          body: formData,
          headers: {
            Authorization: `Bearer ${localStorage.getItem("access_token")}`,
          },
        }
      );

      if (response.ok) {
        const data = await response.json();

        return data.fileUrl;
      } else {
        console.error("Failed to upload file:", response.statusText);
        return null;
      }
    } catch (error) {
      console.error("Error uploading file:", error);
      return null;
    }
  };

  const handleSubmit = async () => {
    const imageUrl = await uploadFile(course.image);
    console.log("Image URL:", imageUrl);
    if (!imageUrl) {
      alert("Image upload failed. Please try again.");
      return;
    }

    const partsWithUploadedMaterials = await Promise.all(
      course.parts.map(async (part) => {
        const materialsWithUrls = await Promise.all(
          part.materials.map(async (material) => {
            if (material.content instanceof File) {
              const fileUrl = await uploadFile(material.content);
              if (!fileUrl) {
                alert(
                  `Material upload failed for ${material.name}. Please try again.`
                );
                return null;
              }
              return { ...material, content: fileUrl };
            } else {
              return material;
            }
          })
        );

        const validMaterials = materialsWithUrls.filter((mat) => mat !== null);

        return { ...part, materials: validMaterials };
      })
    );

    const courseData = {
      name: course.name,
      description: course.description,
      image: imageUrl,
      isFeatured: course.isFeatured,
      parts: partsWithUploadedMaterials,
      tags: course.tags,
    };
    try {
      console.log("Uploading course:", courseData);
      const response = await fetch(
        "https://unique-minds.onrender.com/api/courses/",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${localStorage.getItem("access_token")}`,
          },
          body: JSON.stringify(courseData),
        }
      );

      if (response.ok) {
        const result = await response.json();
        alert("Course uploaded successfully.");
      } else {
        console.error("Failed to upload course:", response.statusText);
      }
    } catch (error) {
      console.error("Error uploading course:", error);
    }
  };

  const handleAddTag = () => {
    if (newTag.trim() === "") {
      alert("Tag cannot be empty.");
      return;
    }
    setCourse((prevCourse) => ({
      ...prevCourse,
      tags: [...prevCourse.tags, newTag.trim()],
    }));
    setNewTag("");
  };

  const handleTagChange = (e) => {
    setNewTag(e.target.value);
  };

  const handleRemoveTag = (tagToRemove) => {
    setCourse((prevCourse) => ({
      ...prevCourse,
      tags: prevCourse.tags.filter((tag) => tag !== tagToRemove),
    }));
  };

  return (
    <div className="bg-customBlue p-8 rounded-lg shadow-lg max-w-6xl mx-auto">
      <h3 className="text-3xl font-bold mb-6 text-white">Upload New Course</h3>

      <div className="flex flex-col lg:flex-row gap-8">
        <div className="bg-white p-6 rounded-lg shadow w-full lg:w-1/3 h-fit">
          <h3 className="text-xl font-semibold mb-4 text-gray-800">
            Course Details
          </h3>
          <div className="flex flex-col space-y-4">
            <input
              type="text"
              name="name"
              value={course.name}
              onChange={handleCourseChange}
              placeholder="Course Name"
              className="p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
            />
            <textarea
              name="description"
              value={course.description}
              onChange={handleCourseChange}
              placeholder="Course Description"
              className="p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500 h-32 resize-none"
            />
            <label className="flex items-center space-x-2">Course Image </label>
            <input
              type="file"
              name="image"
              onChange={handleCourseChange}
              className="p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
            />
          </div>
          {/* Tags Section */}
          <div className="mt-8">
            <h3 className="text-xl font-semibold mb-4 text-gray-800">Tags</h3>
            <div className="flex flex-col space-y-4">
              <div className="flex gap-4 items-center">
                <input
                  type="text"
                  value={newTag}
                  onChange={handleTagChange}
                  placeholder="Add a tag"
                  className="p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
                />
                <button
                  onClick={handleAddTag}
                  className="py-2 px-4 bg-customBlue text-white rounded-lg hover:bg-gray-500 transition duration-300"
                >
                  Add Tag
                </button>
              </div>
              <div className="flex flex-wrap gap-2 mt-4">
                {course.tags.map((tag, index) => (
                  <span
                    key={index}
                    className="bg-gray-200 text-gray-800 px-4 py-2 rounded-full flex items-center"
                  >
                    {tag}
                    <button
                      onClick={() => handleRemoveTag(tag)}
                      className="ml-2 text-red-500"
                    >
                      &times;
                    </button>
                  </span>
                ))}
              </div>
            </div>
          </div>
        </div>

        <div className="flex-1 space-y-8">
          <div className="bg-white p-6 rounded-lg shadow">
            <h3 className="text-xl font-semibold mb-4 text-gray-800">
              Add Part
            </h3>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <input
                type="text"
                name="name"
                value={currentPart.name}
                onChange={handlePartChange}
                placeholder="Part Name"
                className="p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
              />
              <input
                type="number"
                name="sequence"
                value={currentPart.sequence}
                onChange={handlePartChange}
                placeholder="Sequence"
                className="p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
              />
              <textarea
                name="description"
                value={currentPart.description}
                onChange={handlePartChange}
                placeholder="Part Description"
                className="p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500 h-32 resize-none"
              />
            </div>
            <button
              onClick={addPart}
              className="mt-4 py-2 px-4 bg-customBlue text-white rounded-lg hover:bg-gray-500 transition duration-300"
            >
              Add Part
            </button>
          </div>

          <div className="bg-white p-6 rounded-lg shadow">
            <h3 className="text-xl font-semibold mb-4 text-gray-800">
              Add Material
            </h3>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <select
                value={selectedPartSequence || ""}
                onChange={(e) =>
                  setSelectedPartSequence(Number(e.target.value))
                }
                className="p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
              >
                <option value="">Select Part</option>
                {course.parts.map((part) => (
                  <option key={part.sequence} value={part.sequence}>
                    {part.name}
                  </option>
                ))}
              </select>
              <input
                type="text"
                name="name"
                value={currentMaterial.name}
                onChange={handleMaterialChange}
                placeholder="Material Name"
                className="p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
              />
              <select
                name="type"
                value={currentMaterial.type}
                onChange={handleMaterialChange}
                className="p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
              >
                <option value="">Select Type</option>
                <option value="PDF">PDF</option>
                <option value="Video">Video</option>
                <option value="Audio">Audio</option>
              </select>
              <input
                type="file"
                name="content"
                onChange={handleMaterialChange}
                className="p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500"
              />
              <textarea
                name="description"
                value={currentMaterial.description}
                onChange={handleMaterialChange}
                placeholder="Material Description"
                className="p-3 border border-gray-300 rounded-lg focus:outline-none focus:border-blue-500 h-32 resize-none"
              />
            </div>
            <button
              onClick={addMaterial}
              className="mt-4 py-2 px-4 bg-customBlue text-white rounded-lg hover:bg-gray-500 transition duration-300"
            >
              Add Material
            </button>
          </div>
        </div>
      </div>

      <div className="mt-8 flex justify-end">
        <button
          onClick={handleSubmit}
          className="py-2 px-4 bg-white text-customBlue rounded-lg hover:bg-gray-300 transition duration-300"
        >
          Submit Course
        </button>
      </div>
    </div>
  );
};

export default UploadCourse;
