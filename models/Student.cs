namespace Models
{
    public class Student : Model
    {
        public String FirstName;
        public String LastName;
        public String RegistrationNumber;
        public List<Course> Courses;

        public Student(int id, String firstName, String lastName, String RN, List<Course> courses) : base(id)
        {
            FirstName = firstName;
            LastName = lastName;
            RegistrationNumber = RN;
            Courses = courses;
        }
    }
    
}