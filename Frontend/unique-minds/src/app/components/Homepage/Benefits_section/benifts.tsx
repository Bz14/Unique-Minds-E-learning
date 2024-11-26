import React from "react";
import { motion } from "framer-motion";
import {
  CheckCircleIcon,
  UserGroupIcon,
  ClockIcon,
  StarIcon,
} from "@heroicons/react/24/outline";

const benefits = [
  {
    id: 1,
    title: "Personalized Learning",
    description:
      "Tailored lessons to fit each student's unique learning style and pace.",
    icon: <CheckCircleIcon className="h-12 w-12 text-white" />,
  },
  {
    id: 2,
    title: "Quality Instructors",
    description:
      "Learn from qualified professionals with extensive experience in their fields.",
    icon: <UserGroupIcon className="h-12 w-12 text-white" />,
  },
  {
    id: 3,
    title: "Flexible Schedules",
    description:
      "Access courses anytime and anywhere to fit your busy schedule.",
    icon: <ClockIcon className="h-12 w-12 text-white" />,
  },
  {
    id: 4,
    title: "Quality Content",
    description:
      "High-quality, engaging content that makes learning enjoyable and effective.",
    icon: <StarIcon className="h-12 w-12 text-white" />,
  },
];

const BenefitsSection = () => {
  return (
    <div className="py-16 bg-white">
      <motion.div
        className="container mx-auto px-4"
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, amount: 0.3 }}
        transition={{ staggerChildren: 0.2 }}
      >
        <motion.h2
          className="text-3xl font-bold text-center mb-12 text-customBlue"
          initial={{ opacity: 0, y: -50 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 1.5 }}
          viewport={{ once: true, amount: 0.5 }}
        >
          Benefits of Our Platform
        </motion.h2>

        <motion.div
          className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-8"
          initial="hidden"
          animate="visible"
          viewport={{ once: true, amount: 0.3 }}
          transition={{ staggerChildren: 0.2 }}
        >
          {benefits.map((benefit) => (
            <motion.div
              key={benefit.id}
              className="bg-customBlue rounded-lg p-6 shadow-lg flex items-center text-white"
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5, delay: benefit.id * 0.2 }}
              whileHover={{
                scale: 1.05,
                boxShadow: "0 8px 16px rgba(0, 0, 0, 0.2)",
                transition: { duration: 0.3 },
              }}
            >
              <div className="flex-shrink-0 mr-4">{benefit.icon}</div>
              <div>
                <motion.h3
                  className="text-xl font-semibold mb-2"
                  initial={{ opacity: 0, y: 20 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ duration: 0.6 }}
                >
                  {benefit.title}
                </motion.h3>
                <motion.p
                  className="text-white"
                  initial={{ opacity: 0, y: 20 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ duration: 0.6, delay: 0.2 }}
                >
                  {benefit.description}
                </motion.p>
              </div>
            </motion.div>
          ))}
        </motion.div>
      </motion.div>
    </div>
  );
};

export default BenefitsSection;
