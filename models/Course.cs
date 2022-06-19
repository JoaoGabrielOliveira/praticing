namespace Models
{
    public class Course : Model
    {
        public String Code;
        public String Name;

        public Course(int id, String code, String name) : base(id)
        {
            Code = code;
            Name = name;
        }
    }
}