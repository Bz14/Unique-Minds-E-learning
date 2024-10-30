import StudentDashboard from "./student_dashboard/page";
import TeacherDashboard from "./teacher_dashboard/page";

const Dashboard = () => {
  const role = "student";
  return (
    <div>
      {role === "student" ? <StudentDashboard /> : <TeacherDashboard />}
    </div>
  );
};

export default Dashboard;
