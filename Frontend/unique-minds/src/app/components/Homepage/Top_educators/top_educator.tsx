import React, { useState } from "react";
import { motion } from "framer-motion";
import person1 from "@assets/person1.jpg";
import person2 from "@assets/person2.jpg";
import person3 from "@assets/person3.jpg";
import { StaticImageData } from "next/image";

const TopEducators = () => {
  const [educators, setEducators] = useState([
    {
      id: 1,
      name: "Jane Doe",
      image: person1,
      rank: 1,
    },
    {
      id: 2,
      name: "John Smith",
      image: person2,
      rank: 2,
    },
    {
      id: 3,
      name: "Alice Johnson",
      image: person3,
      rank: 3,
    },
  ]);

  return (
    <div className="p-8 bg-gray-100 rounded-lg shadow-md mt-10 mb-10 text-customBlue">
      <motion.h2
        className="text-center text-3xl font-bold mb-8"
        initial={{ opacity: 0, y: -30 }}
        whileInView={{ opacity: 1, y: 0 }}
        viewport={{ once: true, amount: 0.3 }}
        transition={{ duration: 1 }}
      >
        Top Teachers of the Week
      </motion.h2>

      <motion.div
        className="flex justify-center items-end gap-8"
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, amount: 0.3 }}
        transition={{ staggerChildren: 0.2 }}
      >
        {educators.map((educator) => (
          <motion.div
            key={educator.id}
            className={`bg-white rounded-lg shadow-lg overflow-hidden transition-all ${
              educator.rank === 1 ? "scale-110 z-10" : "scale-100"
            }`}
            style={{
              marginTop: educator.rank === 1 ? "-20px" : "0",
            }}
            initial={{ opacity: 0, y: 50 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ duration: 0.8, delay: educator.rank * 0.2 }}
            whileHover={{
              scale: 1.1,
              boxShadow: "0 8px 16px rgba(0, 0, 0, 0.2)",
              y: -10,
              transition: { duration: 0.4 },
            }}
          >
            <motion.img
              src={(educator.image as StaticImageData).src}
              alt={educator.name}
              className="w-48 h-48 object-cover rounded-full mx-auto"
              initial={{ opacity: 0, scale: 0.9 }}
              whileInView={{ opacity: 1, scale: 1 }}
              viewport={{ once: true }}
              transition={{ duration: 0.5 }}
              whileHover={{ scale: 1.05 }}
            />
            <div className="p-4 text-center">
              <motion.h3
                className="text-xl font-semibold text-gray-800"
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                viewport={{ once: true }}
                transition={{ duration: 0.5, delay: 0.2 }}
              >
                {educator.name}
              </motion.h3>
            </div>
          </motion.div>
        ))}
      </motion.div>
    </div>
  );
};

export default TopEducators;
