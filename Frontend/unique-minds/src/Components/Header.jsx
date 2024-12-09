import React from "react";
import { Link } from "react-router-dom";
import { motion } from "framer-motion";
import img1 from "../Assets/img3.jpg";

const Header = () => {
  return (
    <div className="flex items-center justify-center h-screen bg-customBlue text-white">
      <motion.div
        className="flex-1 flex justify-center"
        initial={{ opacity: 0, x: -100 }}
        animate={{ opacity: 1, x: 0, rotate: [0, 10, -10, 0] }}
        transition={{ duration: 1, ease: "easeInOut" }}
      >
        <motion.img
          src={img1}
          alt="E-Learning"
          className="w-3/4 max-w-md rounded-lg shadow-lg"
          whileHover={{ scale: 1.05, rotate: 2 }}
          transition={{ duration: 0.3 }}
        />
      </motion.div>

      <motion.div
        className="flex-1 text-center p-8"
        initial={{ opacity: 0, y: 100 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 1, delay: 0.5, ease: "easeInOut" }}
      >
        <motion.h1
          className="text-5xl font-bold mb-4"
          initial={{ opacity: 0, scale: 0.8 }}
          animate={{ opacity: 1, scale: 1 }}
          transition={{ duration: 0.6, ease: "easeOut" }}
        >
          Welcome to Unique Minds.
        </motion.h1>
        <motion.p
          className="text-lg mb-8"
          initial={{ opacity: 0, x: -50 }}
          animate={{ opacity: 1, x: 0 }}
          transition={{ duration: 0.6, ease: "easeOut", delay: 0.2 }}
        >
          Serving the uniques!
        </motion.p>

        <motion.div
          whileHover={{ scale: 1, rotate: 1 }}
          whileTap={{ scale: 0.95 }}
          transition={{ duration: 0.3 }}
        >
          <Link
            to="/courses"
            className="bg-white text-customBlue px-6 py-3 rounded-lg shadow-lg"
          >
            Browse Courses
          </Link>
        </motion.div>
      </motion.div>
    </div>
  );
};

export default Header;
