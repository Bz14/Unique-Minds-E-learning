"use client";
import img from "../../../assets/img3.jpg";
import { motion } from "framer-motion";
import Link from "next/link";

const HeroSection = () => {
  return (
    <div
      className="relative w-full h-screen bg-cover bg-center"
      style={{ backgroundImage: `url(${img.src})` }}
    >
      <div className="absolute inset-0 bg-black opacity-50"></div>
      <motion.div
        className="relative z-10 flex flex-col items-center justify-center h-full text-center text-white"
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
            href="/courses"
            className="bg-white text-customBlue px-6 py-3 rounded-lg shadow-lg"
          >
            Browse Courses
          </Link>
        </motion.div>
      </motion.div>
    </div>
  );
};

export default HeroSection;
