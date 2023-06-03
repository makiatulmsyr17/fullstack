import Navbar from "../components/Navbar";
import axios from "axios";
import { useForm, useFormState } from "react-hook-form";
import { useNavigate, Link, useParams } from "react-router-dom";
import { useEffect } from "react";


const EditData = () => {
    const { register, handleSubmit,
        setValue,
        formState: { errors }, } = useForm({
            defaultValues: {
                nama: "",
                kelas: "",
                prodi: "",
            },
            mode: "onBlur",
        });

    const navigate = useNavigate();
    const { id } = useParams();
    useEffect(() => {
        getDataId();
    }, []);
    const getDataId = async () => {
        try {
            const response = await axios.get(`http://localhost:9000/read/${id}`);
            console.log(response.data.data);
            const res = await response.data.data;
            setValue("nama", res.nama);
            setValue("kelas", res.kelas);
            setValue("prodi", res.prodi);


        } catch (error) {
            console.error(error);
        }
    }

    const editSubmit = async (data) => {
        try {
            const response = await axios.put(`http://localhost:9000/update/${id}`,
                data,
                {
                    headers: {
                        "Content-Type": "application/json",
                    },
                    withCredentials: false,
                }
            );
            console.log(response);
            navigate("/view")

        } catch (error) {
            console.error(error);
        }
    };

    return (
        <>
            <Navbar />
            <div className="flex justify-center items-center h-screen">
                <form onSubmit={handleSubmit(editSubmit)} className="flex flex-col items-center justify-center w-[50vw] h-[80vh] shadow-md border-2 text-lg gap-4">
                    <h1 className="my-6">edit Data</h1>
                    {/*nama*/}
                    <div className="flex flex-col gap-1">
                        <label htmlFor="nama">nama</label>
                        <input type="text" id="nama" placeholder="masukan nama " className="p-4 bg-slate-100 rounded-md w-[40vw] border-2 border-slate-200"  {...register("nama", { required: "nama harus diisi" })} />
                        {errors.nama && (
                            <span className="text-red-600">(errors.nama.message) </span>)}
                    </div>
                    {/*kelas*/}
                    <div className="flex flex-col">
                        <label htmlFor="kelas">kelas</label>
                        <input type="text" id="kelas" placeholder="masukan kelas " className="p-4 bg-slate-100 rounded-md w-[40vw] border-2 border-slate-200"  {...register("Kelas", { required: "kelas harus diisi" })} />
                        {errors.kelas && (
                            <span className="text-red-600">(errors.nama.message) </span>)}
                    </div>
                    {/*prodi*/}
                    <div className="flex flex-col">
                        <label htmlFor="prodi">prodi</label>
                        <input type="text" id="prodi" placeholder="masukan prodi " className="p-4 bg-slate-100 rounded-md w-[40vw] border-2 border-slate-200"  {...register("Prodi", { required: "prodi harus diisi" })} />
                        {errors.prodi && (
                            <span className="text-red-600">(errors.prodi.message) </span>)}
                    </div>
                    <button className=" w-[10vw] h-[5vh] mt-10 bg-blue-500 hover:bg-blue-400 text-white forn-semibold">add</button>
                </form>
            </div>
        </>
    );
};
export default EditData;