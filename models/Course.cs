namespace Models
{
    public class Course : Model<long>
    {
        public String Code  {get; set;}
        public String Name {get; set;}

        public Course(int id, String code, String name) : base(id)
        {
            Code = code;
            Name = name;
        }
    }
}