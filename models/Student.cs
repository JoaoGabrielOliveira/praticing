namespace Models
{
    public class Student : Model<long>
    {
        public String FirstName { get; set; }
        public String LastName { get; set; }
        public String RegistrationNumber { get; set; }
        public List<Course> Courses { get; set; }

        public Student(int id, String firstName, String lastName, String RN, List<Course> courses) : base(id)
        {
            FirstName = firstName;
            LastName = lastName;
            RegistrationNumber = RN;
            Courses = courses;
        }

        public Student() : base() {}
    }
    
}