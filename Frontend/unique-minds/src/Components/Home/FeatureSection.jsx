import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { motion } from "framer-motion";

// FeaturedCourses component
const FeaturedCourses = () => {
  const [courses, setCourses] = useState([]);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  if (loading) return <p className="text-center py-16">Loading...</p>;
  if (error) return <p className="text-center py-16 text-red-500">{error}</p>;

  const containerVariants = {
    hidden: { opacity: 0, y: 100 },
    visible: {
      opacity: 1,
      y: 0,
      transition: {
        delayChildren: 0.5,
        staggerChildren: 0.2,
      },
    },
  };

  const courseVariants = {
    hidden: { opacity: 0, y: 50 },
    visible: { opacity: 1, y: 0 },
  };

  return (
    <div className="py-16 bg-gray-100">
      <motion.div
        className="container mx-auto px-4"
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, amount: 0.3 }}
        variants={containerVariants}
      >
        <motion.h2
          className="text-3xl font-bold text-center mb-12 text-customBlue"
          initial={{ opacity: 0, y: -50 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ duration: 1.5 }}
          viewport={{ once: true, amount: 0.5 }}
        >
          Featured Courses
        </motion.h2>

        <motion.div
          className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8"
          initial="hidden"
          whileInView="visible"
          variants={containerVariants}
          viewport={{ once: true, amount: 0.3 }}
        >
          {courses != null &&
            courses.length > 0 &&
            courses.map((course) => (
              <motion.div
                key={course._id}
                className="bg-white rounded-lg shadow-lg overflow-hidden"
                variants={courseVariants}
                transition={{ duration: 0.8 }}
                whileHover={{ scale: 1.02 }}
              >
                <motion.img
                  src={course.image}
                  alt={course.name}
                  className="w-full h-40 object-cover"
                  initial={{ scale: 1.05 }}
                  whileHover={{ scale: 1.1 }}
                  transition={{ duration: 0.4 }}
                />
                <div className="p-6">
                  <motion.h3
                    className="text-xl font-semibold mb-2"
                    initial={{ opacity: 0, y: 20 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ duration: 0.6 }}
                  >
                    {course.name}
                  </motion.h3>
                  <motion.p
                    className="text-gray-700 mb-4"
                    initial={{ opacity: 0, y: 20 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ duration: 0.6, delay: 0.2 }}
                  >
                    {course.description}
                  </motion.p>
                  <motion.div
                    whileHover={{ scale: 1.05 }}
                    whileTap={{ scale: 0.95 }}
                    transition={{ duration: 0.3 }}
                  >
                    <Link
                      to={`/courses/${course._id}`}
                      className="bg-customBlue text-white px-4 py-2 rounded-lg hover:bg-gray-600 transition"
                    >
                      Learn More
                    </Link>
                  </motion.div>
                </div>
              </motion.div>
            ))}
        </motion.div>
      </motion.div>
    </div>
  );
};

export default FeaturedCourses;
