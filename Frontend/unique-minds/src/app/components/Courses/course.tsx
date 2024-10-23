import Image, { StaticImageData } from "next/image";

interface CourseProps {
  course: {
    image: StaticImageData;
    title: string;
    description: string;
  };
}

const Course = (props: CourseProps) => {
  return (
    <div className="course bg-white shadow-lg rounded-lg overflow-hidden transform transition hover:scale-105 hover:shadow-xl">
      <div className="course__img">
        <Image
          src={props.course.image}
          alt={props.course.title}
          width={500}
          height={300}
          className="object-cover w-full h-[300px] rounded-t-lg"
        />
      </div>
      <div className="p-6">
        <h3 className="text-xl font-bold text-gray-800 mb-2">
          {props.course.title}
        </h3>
        <p className="text-gray-600 mb-4">{props.course.description}</p>
        <button className="px-4 py-2 bg-customBlue text-white font-semibold rounded hover:bg-gray-400 transition-colors duration-300">
          View Course
        </button>
      </div>
    </div>
  );
};

export default Course;
