package common

// PeriodicTable represents the chemical periodic table.  It defines
// the necessary constants for chemical elements, their atomic
// numbers, atomic weights, isotopes and symbol string
// representations.  It provides a set of convenience methods to
// access and utilise the said data.
var PeriodicTable = map[string]Element{
	"NONE":   Element{0, "NONE", "NONE", 0.0, -1, []int8{}},
	"H":      Element{1, "H", "Hydrogen", 1.008, 1, []int8{}},
	"He":     Element{2, "He", "Helium", 4.003, 0, []int8{}},
	"Li":     Element{3, "Li", "Lithium", 6.941, 1, []int8{}},
	"Be":     Element{4, "Be", "Beryllium", 9.012, 2, []int8{}},
	"B":      Element{5, "B", "Boron", 10.812, 3, []int8{}},
	"C":      Element{6, "C", "Carbon", 12.011, 4, []int8{}},
	"N":      Element{7, "N", "Nitrogen", 14.007, 3, []int8{}},
	"O":      Element{8, "O", "Oxygen", 15.999, 2, []int8{}},
	"F":      Element{9, "F", "Fluorine", 18.998, 1, []int8{}},
	"Ne":     Element{10, "Ne", "Neon", 20.18, 0, []int8{}},
	"Na":     Element{11, "Na", "Sodium", 22.99, 1, []int8{}},
	"Mg":     Element{12, "Mg", "Magnesium", 24.305, 2, []int8{}},
	"Al":     Element{13, "Al", "Aluminium", 26.982, 6, []int8{}},
	"Si":     Element{14, "Si", "Silicon", 28.086, 4, []int8{}},
	"P":      Element{15, "P", "Phosphorus", 30.974, 3, []int8{}},
	"S":      Element{16, "S", "Sulfur", 32.067, 6, []int8{}},
	"Cl":     Element{17, "Cl", "Chlorine", 35.453, 1, []int8{}},
	"Ar":     Element{18, "Ar", "Argon", 39.948, 0, []int8{}},
	"K":      Element{19, "K", "Potassium", 39.098, 1, []int8{}},
	"Ca":     Element{20, "Ca", "Calcium", 40.078, 2, []int8{}},
	"Sc":     Element{21, "Sc", "Scandium", 44.956, -1, []int8{}},
	"Ti":     Element{22, "Ti", "Titanium", 47.867, -1, []int8{}},
	"V":      Element{23, "V", "Vanadium", 50.942, -1, []int8{}},
	"Cr":     Element{24, "Cr", "Chromium", 51.996, -1, []int8{}},
	"Mn":     Element{25, "Mn", "Manganese", 54.938, -1, []int8{}},
	"Fe":     Element{26, "Fe", "Iron", 55.845, -1, []int8{}},
	"Co":     Element{27, "Co", "Cobalt", 58.933, -1, []int8{}},
	"Ni":     Element{28, "Ni", "Nickel", 58.693, -1, []int8{}},
	"Cu":     Element{29, "Cu", "Copper", 63.546, -1, []int8{}},
	"Zn":     Element{30, "Zn", "Zinc", 65.39, -1, []int8{}},
	"Ga":     Element{31, "Ga", "Gallium", 69.723, 3, []int8{}},
	"Ge":     Element{32, "Ge", "Germanium", 72.61, 4, []int8{}},
	"As":     Element{33, "As", "Arsenic", 74.922, 3, []int8{}},
	"Se":     Element{34, "Se", "Selenium", 78.96, 2, []int8{}},
	"Br":     Element{35, "Br", "Bromine", 79.904, 1, []int8{}},
	"Kr":     Element{36, "Kr", "Krypton", 83.8, 0, []int8{}},
	"Rb":     Element{37, "Rb", "Rubidium", 85.468, 1, []int8{}},
	"Sr":     Element{38, "Sr", "Strontium", 87.62, 2, []int8{}},
	"Y":      Element{39, "Y", "Yttrium", 88.906, -1, []int8{}},
	"Zr":     Element{40, "Zr", "Zirconium", 91.224, -1, []int8{}},
	"Nb":     Element{41, "Nb", "Niobium", 92.906, -1, []int8{}},
	"Mo":     Element{42, "Mo", "Molybdenum", 95.94, -1, []int8{}},
	"Tc":     Element{43, "Tc", "Technetium", 98, -1, []int8{}},
	"Ru":     Element{44, "Ru", "Ruthenium", 101.07, -1, []int8{}},
	"Rh":     Element{45, "Rh", "Rhodium", 102.906, -1, []int8{}},
	"Pd":     Element{46, "Pd", "Palladium", 106.42, -1, []int8{}},
	"Ag":     Element{47, "Ag", "Silver", 107.868, -1, []int8{}},
	"Cd":     Element{48, "Cd", "Cadmium", 112.412, -1, []int8{}},
	"In":     Element{49, "In", "Indium", 114.818, 3, []int8{}},
	"Sn":     Element{50, "Sn", "Tin", 118.711, 4, []int8{}},
	"Sb":     Element{51, "Sb", "Antimony", 121.76, 3, []int8{}},
	"Te":     Element{52, "Te", "Tellurium", 127.6, 2, []int8{}},
	"I":      Element{53, "I", "Iodine", 126.904, 1, []int8{}},
	"Xe":     Element{54, "Xe", "Xenon", 131.29, 0, []int8{}},
	"Cs":     Element{55, "Cs", "Caesium", 132.905, 1, []int8{}},
	"Ba":     Element{56, "Ba", "Barium", 137.328, 2, []int8{}},
	"La":     Element{57, "La", "Lanthanum", 138.906, -1, []int8{}},
	"Ce":     Element{58, "Ce", "Cerium", 140.116, -1, []int8{}},
	"Pr":     Element{59, "Pr", "Praseodymium", 140.908, -1, []int8{}},
	"Nd":     Element{60, "Nd", "Neodymium", 144.24, -1, []int8{}},
	"Pm":     Element{61, "Pm", "Promethium", 145, -1, []int8{}},
	"Sm":     Element{62, "Sm", "Samarium", 150.36, -1, []int8{}},
	"Eu":     Element{63, "Eu", "Europium", 151.964, -1, []int8{}},
	"Gd":     Element{64, "Gd", "Gadolinium", 157.25, -1, []int8{}},
	"Tb":     Element{65, "Tb", "Terbium", 158.925, -1, []int8{}},
	"Dy":     Element{66, "Dy", "Dysprosium", 162.5, -1, []int8{}},
	"Ho":     Element{67, "Ho", "Holmium", 164.93, -1, []int8{}},
	"Er":     Element{68, "Er", "Erbium", 167.26, -1, []int8{}},
	"Tm":     Element{69, "Tm", "Thulium", 168.934, -1, []int8{}},
	"Yb":     Element{70, "Yb", "Ytterbium", 173.04, -1, []int8{}},
	"Lu":     Element{71, "Lu", "Lutetium", 174.967, -1, []int8{}},
	"Hf":     Element{72, "Hf", "Hafnium", 178.49, -1, []int8{}},
	"Ta":     Element{73, "Ta", "Tantalum", 180.948, -1, []int8{}},
	"W":      Element{74, "W", "Tungsten", 183.84, -1, []int8{}},
	"Re":     Element{75, "Re", "Rhenium", 186.207, -1, []int8{}},
	"Os":     Element{76, "Os", "Osmium", 190.23, -1, []int8{}},
	"Ir":     Element{77, "Ir", "Iridium", 192.217, -1, []int8{}},
	"Pt":     Element{78, "Pt", "Platinum", 195.078, -1, []int8{}},
	"Au":     Element{79, "Au", "Gold", 196.967, -1, []int8{}},
	"Hg":     Element{80, "Hg", "Mercury", 200.59, -1, []int8{}},
	"Tl":     Element{81, "Tl", "Thallium", 204.383, 3, []int8{}},
	"Pb":     Element{82, "Pb", "Lead", 207.2, 4, []int8{}},
	"Bi":     Element{83, "Bi", "Bismuth", 208.98, 3, []int8{}},
	"Po":     Element{84, "Po", "Polonium", 209, 2, []int8{}},
	"At":     Element{85, "At", "Astatine", 210, 1, []int8{}},
	"Rn":     Element{86, "Rn", "Radon", 222, 0, []int8{}},
	"Fr":     Element{87, "Fr", "Francium", 223, 1, []int8{}},
	"Ra":     Element{88, "Ra", "Radium", 226, 2, []int8{}},
	"Ac":     Element{89, "Ac", "Actinium", 227, -1, []int8{}},
	"Th":     Element{90, "Th", "Thorium", 232.038, -1, []int8{}},
	"Pa":     Element{91, "Pa", "Protactinium", 231.036, -1, []int8{}},
	"U":      Element{92, "U", "Uranium", 238.029, -1, []int8{}},
	"Np":     Element{93, "Np", "Neptunium", 237, -1, []int8{}},
	"Pu":     Element{94, "Pu", "Plutonium", 244, -1, []int8{}},
	"Am":     Element{95, "Am", "Americium", 243, -1, []int8{}},
	"Cm":     Element{96, "Cm", "Curium", 247, -1, []int8{}},
	"Bk":     Element{97, "Bk", "Berkelium", 247, -1, []int8{}},
	"Cf":     Element{98, "Cf", "Californium", 251, -1, []int8{}},
	"Es":     Element{99, "Es", "Einsteinium", 252, -1, []int8{}},
	"Fm":     Element{100, "Fm", "Fermium", 257, -1, []int8{}},
	"Md":     Element{101, "Md", "Mendelevium", 258, -1, []int8{}},
	"No":     Element{102, "No", "Nobelium", 259, -1, []int8{}},
	"Lr":     Element{103, "Lr", "Lawrencium", 262, -1, []int8{}},
	"Rf":     Element{104, "Rf", "Rutherfordium", 267, -1, []int8{}},
	"Db":     Element{105, "Db", "Dubnium", 268, -1, []int8{}},
	"Sg":     Element{106, "Sg", "Seaborgium", 269, -1, []int8{}},
	"Bh":     Element{107, "Bh", "Bohrium", 270, -1, []int8{}},
	"Hs":     Element{108, "Hs", "Hassium", 269, -1, []int8{}},
	"Mt":     Element{109, "Mt", "Meitnerium", 278, -1, []int8{}},
	"Ds":     Element{110, "Ds", "Darmstadtium", 281, -1, []int8{}},
	"Rg":     Element{111, "Rg", "Roentgenium", 281, -1, []int8{}},
	"Cn":     Element{112, "Cn", "Copernicium", 285, -1, []int8{}},
	"Uut":    Element{113, "Uut", "Ununtrium", 286, -1, []int8{}},
	"Fl":     Element{114, "Fl", "Flerovium", 289, -1, []int8{}},
	"Uup":    Element{115, "Uup", "Ununpentium", 288, -1, []int8{}},
	"Lv":     Element{116, "Lv", "Livermorium", 293, -1, []int8{}},
	"Uus":    Element{117, "Uus", "Ununseptium", 294, -1, []int8{}},
	"Uuo":    Element{118, "Uuo", "Ununoctium", 294, -1, []int8{}},
	"D":      Element{1, "D", "", 2.014101778, 1, []int8{}},
	"H_2":    Element{1, "H_2", "", 2.014101778, 1, []int8{}},
	"T":      Element{1, "T", "", 3.016049278, 1, []int8{}},
	"H_3":    Element{1, "H_3", "", 3.016049278, 1, []int8{}},
	"H_4":    Element{1, "H_4", "", 4.02781, 1, []int8{}},
	"H_5":    Element{1, "H_5", "", 5.03531, 1, []int8{}},
	"H_6":    Element{1, "H_6", "", 6.04494, 1, []int8{}},
	"H_7":    Element{1, "H_7", "", 7.05275, 1, []int8{}},
	"C_8":    Element{6, "C_8", "", 8.037675, 4, []int8{}},
	"C_9":    Element{6, "C_9", "", 9.0310367, 4, []int8{}},
	"C_10":   Element{6, "C_10", "", 10.0168532, 4, []int8{}},
	"C_11":   Element{6, "C_11", "", 11.0114336, 4, []int8{}},
	"C_13":   Element{6, "C_13", "", 13.00335484, 4, []int8{}},
	"C_14":   Element{6, "C_14", "", 14.00324199, 4, []int8{}},
	"C_15":   Element{6, "C_15", "", 15.0105993, 4, []int8{}},
	"C_16":   Element{6, "C_16", "", 16.014701, 4, []int8{}},
	"C_17":   Element{6, "C_17", "", 17.022586, 4, []int8{}},
	"C_18":   Element{6, "C_18", "", 18.02676, 4, []int8{}},
	"C_19":   Element{6, "C_19", "", 19.03481, 4, []int8{}},
	"C_20":   Element{6, "C_20", "", 20.04032, 4, []int8{}},
	"C_21":   Element{6, "C_21", "", 21.04934, 4, []int8{}},
	"C_22":   Element{6, "C_22", "", 22.0572, 4, []int8{}},
	"N_10":   Element{7, "N_10", "", 10.04165, 3, []int8{}},
	"N_11":   Element{7, "N_11", "", 11.02609, 3, []int8{}},
	"N_12":   Element{7, "N_12", "", 12.0186132, 3, []int8{}},
	"N_13":   Element{7, "N_13", "", 13.00573861, 3, []int8{}},
	"N_15":   Element{7, "N_15", "", 15.0001089, 3, []int8{}},
	"N_16":   Element{7, "N_16", "", 16.0061017, 3, []int8{}},
	"N_17":   Element{7, "N_17", "", 17.00845, 3, []int8{}},
	"N_18":   Element{7, "N_18", "", 18.014079, 3, []int8{}},
	"N_19":   Element{7, "N_19", "", 19.017029, 3, []int8{}},
	"N_20":   Element{7, "N_20", "", 20.02337, 3, []int8{}},
	"N_21":   Element{7, "N_21", "", 21.02711, 3, []int8{}},
	"N_22":   Element{7, "N_22", "", 22.03439, 3, []int8{}},
	"N_23":   Element{7, "N_23", "", 23.04122, 3, []int8{}},
	"N_24":   Element{7, "N_24", "", 24.05104, 3, []int8{}},
	"N_25":   Element{7, "N_25", "", 25.06066, 3, []int8{}},
	"O_12":   Element{8, "O_12", "", 12.034405, 2, []int8{}},
	"O_13":   Element{8, "O_13", "", 13.024812, 2, []int8{}},
	"O_14":   Element{8, "O_14", "", 14.00859625, 2, []int8{}},
	"O_15":   Element{8, "O_15", "", 15.0030656, 2, []int8{}},
	"O_17":   Element{8, "O_17", "", 16.9991317, 2, []int8{}},
	"O_18":   Element{8, "O_18", "", 17.999161, 2, []int8{}},
	"O_19":   Element{8, "O_19", "", 19.00358, 2, []int8{}},
	"O_20":   Element{8, "O_20", "", 20.0040767, 2, []int8{}},
	"O_21":   Element{8, "O_21", "", 21.008656, 2, []int8{}},
	"O_22":   Element{8, "O_22", "", 22.00997, 2, []int8{}},
	"O_23":   Element{8, "O_23", "", 23.01569, 2, []int8{}},
	"O_24":   Element{8, "O_24", "", 24.02047, 2, []int8{}},
	"O_25":   Element{8, "O_25", "", 25.02946, 2, []int8{}},
	"O_26":   Element{8, "O_26", "", 26.03834, 2, []int8{}},
	"O_27":   Element{8, "O_27", "", 27.04826, 2, []int8{}},
	"O_28":   Element{8, "O_28", "", 28.05781, 2, []int8{}},
	"F_14":   Element{9, "F_14", "", 14.03506, 1, []int8{}},
	"F_15":   Element{9, "F_15", "", 15.01801, 1, []int8{}},
	"F_16":   Element{9, "F_16", "", 16.011466, 1, []int8{}},
	"F_17":   Element{9, "F_17", "", 17.00209524, 1, []int8{}},
	"F_18":   Element{9, "F_18", "", 18.000938, 1, []int8{}},
	"F_20":   Element{9, "F_20", "", 19.99998132, 1, []int8{}},
	"F_21":   Element{9, "F_21", "", 20.999949, 1, []int8{}},
	"F_22":   Element{9, "F_22", "", 22.002999, 1, []int8{}},
	"F_23":   Element{9, "F_23", "", 23.00357, 1, []int8{}},
	"F_24":   Element{9, "F_24", "", 24.00812, 1, []int8{}},
	"F_25":   Element{9, "F_25", "", 25.0121, 1, []int8{}},
	"F_26":   Element{9, "F_26", "", 26.01962, 1, []int8{}},
	"F_27":   Element{9, "F_27", "", 27.02676, 1, []int8{}},
	"F_28":   Element{9, "F_28", "", 28.03567, 1, []int8{}},
	"F_29":   Element{9, "F_29", "", 29.04326, 1, []int8{}},
	"F_30":   Element{9, "F_30", "", 30.0525, 1, []int8{}},
	"F_31":   Element{9, "F_31", "", 31.06043, 1, []int8{}},
	"P_24":   Element{15, "P_24", "", 24.03435, 3, []int8{}},
	"P_25":   Element{15, "P_25", "", 25.02026, 3, []int8{}},
	"P_26":   Element{15, "P_26", "", 26.01178, 3, []int8{}},
	"P_27":   Element{15, "P_27", "", 26.99923, 3, []int8{}},
	"P_28":   Element{15, "P_28", "", 27.992315, 3, []int8{}},
	"P_29":   Element{15, "P_29", "", 28.9818006, 3, []int8{}},
	"P_30":   Element{15, "P_30", "", 29.9783138, 3, []int8{}},
	"P_32":   Element{15, "P_32", "", 31.97390727, 3, []int8{}},
	"P_33":   Element{15, "P_33", "", 32.9717255, 3, []int8{}},
	"P_34":   Element{15, "P_34", "", 33.973636, 3, []int8{}},
	"P_35":   Element{15, "P_35", "", 34.9733141, 3, []int8{}},
	"P_36":   Element{15, "P_36", "", 35.97826, 3, []int8{}},
	"P_37":   Element{15, "P_37", "", 36.97961, 3, []int8{}},
	"P_38":   Element{15, "P_38", "", 37.98416, 3, []int8{}},
	"P_39":   Element{15, "P_39", "", 38.98618, 3, []int8{}},
	"P_40":   Element{15, "P_40", "", 39.9913, 3, []int8{}},
	"P_41":   Element{15, "P_41", "", 40.99434, 3, []int8{}},
	"P_42":   Element{15, "P_42", "", 42.00101, 3, []int8{}},
	"P_43":   Element{15, "P_43", "", 43.00619, 3, []int8{}},
	"P_44":   Element{15, "P_44", "", 44.01299, 3, []int8{}},
	"P_45":   Element{15, "P_45", "", 45.01922, 3, []int8{}},
	"P_46":   Element{15, "P_46", "", 46.02738, 3, []int8{}},
	"S_26":   Element{16, "S_26", "", 26.02788, 2, []int8{}},
	"S_27":   Element{16, "S_27", "", 27.01883, 2, []int8{}},
	"S_28":   Element{16, "S_28", "", 28.00437, 2, []int8{}},
	"S_29":   Element{16, "S_29", "", 28.99661, 2, []int8{}},
	"S_30":   Element{16, "S_30", "", 29.984903, 2, []int8{}},
	"S_31":   Element{16, "S_31", "", 30.9795547, 2, []int8{}},
	"S_33":   Element{16, "S_33", "", 32.97145876, 2, []int8{}},
	"S_34":   Element{16, "S_34", "", 33.9678669, 2, []int8{}},
	"S_35":   Element{16, "S_35", "", 34.96903216, 2, []int8{}},
	"S_36":   Element{16, "S_36", "", 35.96708076, 2, []int8{}},
	"S_37":   Element{16, "S_37", "", 36.97112557, 2, []int8{}},
	"S_38":   Element{16, "S_38", "", 37.971163, 2, []int8{}},
	"S_39":   Element{16, "S_39", "", 38.97513, 2, []int8{}},
	"S_40":   Element{16, "S_40", "", 39.97545, 2, []int8{}},
	"S_41":   Element{16, "S_41", "", 40.97958, 2, []int8{}},
	"S_42":   Element{16, "S_42", "", 41.98102, 2, []int8{}},
	"S_43":   Element{16, "S_43", "", 42.98715, 2, []int8{}},
	"S_44":   Element{16, "S_44", "", 43.99021, 2, []int8{}},
	"S_45":   Element{16, "S_45", "", 44.99651, 2, []int8{}},
	"S_46":   Element{16, "S_46", "", 46.00075, 2, []int8{}},
	"S_47":   Element{16, "S_47", "", 47.00859, 2, []int8{}},
	"S_48":   Element{16, "S_48", "", 48.01417, 2, []int8{}},
	"S_49":   Element{16, "S_49", "", 49.02362, 2, []int8{}},
	"Cl_28":  Element{17, "Cl_28", "", 28.02851, 1, []int8{}},
	"Cl_29":  Element{17, "Cl_29", "", 29.01411, 1, []int8{}},
	"Cl_30":  Element{17, "Cl_30", "", 30.00477, 1, []int8{}},
	"Cl_31":  Element{17, "Cl_31", "", 30.99241, 1, []int8{}},
	"Cl_32":  Element{17, "Cl_32", "", 31.98569, 1, []int8{}},
	"Cl_33":  Element{17, "Cl_33", "", 32.9774519, 1, []int8{}},
	"Cl_34":  Element{17, "Cl_34", "", 33.97376282, 1, []int8{}},
	"Cl_36":  Element{17, "Cl_36", "", 35.96830698, 1, []int8{}},
	"Cl_37":  Element{17, "Cl_37", "", 36.96590259, 1, []int8{}},
	"Cl_38":  Element{17, "Cl_38", "", 37.96801043, 1, []int8{}},
	"Cl_39":  Element{17, "Cl_39", "", 38.9680082, 1, []int8{}},
	"Cl_40":  Element{17, "Cl_40", "", 39.97042, 1, []int8{}},
	"Cl_41":  Element{17, "Cl_41", "", 40.97068, 1, []int8{}},
	"Cl_42":  Element{17, "Cl_42", "", 41.97325, 1, []int8{}},
	"Cl_43":  Element{17, "Cl_43", "", 42.97405, 1, []int8{}},
	"Cl_44":  Element{17, "Cl_44", "", 43.97828, 1, []int8{}},
	"Cl_45":  Element{17, "Cl_45", "", 44.98029, 1, []int8{}},
	"Cl_46":  Element{17, "Cl_46", "", 45.98421, 1, []int8{}},
	"Cl_47":  Element{17, "Cl_47", "", 46.98871, 1, []int8{}},
	"Cl_48":  Element{17, "Cl_48", "", 47.99495, 1, []int8{}},
	"Cl_49":  Element{17, "Cl_49", "", 49.00032, 1, []int8{}},
	"Cl_50":  Element{17, "Cl_50", "", 50.00784, 1, []int8{}},
	"Cl_51":  Element{17, "Cl_51", "", 51.01449, 1, []int8{}},
	"Br_67":  Element{35, "Br_67", "", 66.96479, 1, []int8{}},
	"Br_68":  Element{35, "Br_68", "", 67.95852, 1, []int8{}},
	"Br_69":  Element{35, "Br_69", "", 68.95011, 1, []int8{}},
	"Br_70":  Element{35, "Br_70", "", 69.94479, 1, []int8{}},
	"Br_71":  Element{35, "Br_71", "", 70.93874, 1, []int8{}},
	"Br_72":  Element{35, "Br_72", "", 71.93664, 1, []int8{}},
	"Br_73":  Element{35, "Br_73", "", 72.93169, 1, []int8{}},
	"Br_74":  Element{35, "Br_74", "", 73.929891, 1, []int8{}},
	"Br_75":  Element{35, "Br_75", "", 74.925776, 1, []int8{}},
	"Br_76":  Element{35, "Br_76", "", 75.924541, 1, []int8{}},
	"Br_77":  Element{35, "Br_77", "", 76.921379, 1, []int8{}},
	"Br_78":  Element{35, "Br_78", "", 77.921146, 1, []int8{}},
	"Br_79":  Element{35, "Br_79", "", 78.9183371, 1, []int8{}},
	"Br_81":  Element{35, "Br_81", "", 80.9162906, 1, []int8{}},
	"Br_82":  Element{35, "Br_82", "", 81.9168041, 1, []int8{}},
	"Br_83":  Element{35, "Br_83", "", 82.91518, 1, []int8{}},
	"Br_84":  Element{35, "Br_84", "", 83.916479, 1, []int8{}},
	"Br_85":  Element{35, "Br_85", "", 84.915608, 1, []int8{}},
	"Br_86":  Element{35, "Br_86", "", 85.918798, 1, []int8{}},
	"Br_87":  Element{35, "Br_87", "", 86.920711, 1, []int8{}},
	"Br_88":  Element{35, "Br_88", "", 87.92407, 1, []int8{}},
	"Br_89":  Element{35, "Br_89", "", 88.92639, 1, []int8{}},
	"Br_90":  Element{35, "Br_90", "", 89.93063, 1, []int8{}},
	"Br_91":  Element{35, "Br_91", "", 90.93397, 1, []int8{}},
	"Br_92":  Element{35, "Br_92", "", 91.93926, 1, []int8{}},
	"Br_93":  Element{35, "Br_93", "", 92.94305, 1, []int8{}},
	"Br_94":  Element{35, "Br_94", "", 93.94868, 1, []int8{}},
	"Br_95":  Element{35, "Br_95", "", 94.95287, 1, []int8{}},
	"Br_96":  Element{35, "Br_96", "", 95.95853, 1, []int8{}},
	"Br_97":  Element{35, "Br_97", "", 96.9628, 1, []int8{}},
	"I_108":  Element{53, "I_108", "", 107.94348, 1, []int8{}},
	"I_109":  Element{53, "I_109", "", 108.93815, 1, []int8{}},
	"I_110":  Element{53, "I_110", "", 109.93524, 1, []int8{}},
	"I_111":  Element{53, "I_111", "", 110.93028, 1, []int8{}},
	"I_112":  Element{53, "I_112", "", 111.92797, 1, []int8{}},
	"I_113":  Element{53, "I_113", "", 112.92364, 1, []int8{}},
	"I_114":  Element{53, "I_114", "", 113.92185, 1, []int8{}},
	"I_115":  Element{53, "I_115", "", 114.91805, 1, []int8{}},
	"I_116":  Element{53, "I_116", "", 115.91681, 1, []int8{}},
	"I_117":  Element{53, "I_117", "", 116.91365, 1, []int8{}},
	"I_118":  Element{53, "I_118", "", 117.913074, 1, []int8{}},
	"I_119":  Element{53, "I_119", "", 118.91007, 1, []int8{}},
	"I_120":  Element{53, "I_120", "", 119.910048, 1, []int8{}},
	"I_121":  Element{53, "I_121", "", 120.907367, 1, []int8{}},
	"I_122":  Element{53, "I_122", "", 121.907589, 1, []int8{}},
	"I_123":  Element{53, "I_123", "", 122.905589, 1, []int8{}},
	"I_124":  Element{53, "I_124", "", 123.9062099, 1, []int8{}},
	"I_125":  Element{53, "I_125", "", 124.9046302, 1, []int8{}},
	"I_126":  Element{53, "I_126", "", 125.905624, 1, []int8{}},
	"I_128":  Element{53, "I_128", "", 127.905809, 1, []int8{}},
	"I_129":  Element{53, "I_129", "", 128.904988, 1, []int8{}},
	"I_130":  Element{53, "I_130", "", 129.906674, 1, []int8{}},
	"I_131":  Element{53, "I_131", "", 130.9061246, 1, []int8{}},
	"I_132":  Element{53, "I_132", "", 131.907997, 1, []int8{}},
	"I_133":  Element{53, "I_133", "", 132.907797, 1, []int8{}},
	"I_134":  Element{53, "I_134", "", 133.909744, 1, []int8{}},
	"I_135":  Element{53, "I_135", "", 134.910048, 1, []int8{}},
	"I_136":  Element{53, "I_136", "", 135.91465, 1, []int8{}},
	"I_137":  Element{53, "I_137", "", 136.917871, 1, []int8{}},
	"I_138":  Element{53, "I_138", "", 137.92235, 1, []int8{}},
	"I_139":  Element{53, "I_139", "", 138.9261, 1, []int8{}},
	"I_140":  Element{53, "I_140", "", 139.931, 1, []int8{}},
	"I_141":  Element{53, "I_141", "", 140.93503, 1, []int8{}},
	"I_142":  Element{53, "I_142", "", 141.94018, 1, []int8{}},
	"I_143":  Element{53, "I_143", "", 142.94456, 1, []int8{}},
	"I_144":  Element{53, "I_144", "", 143.94999, 1, []int8{}},
	"Q_STAR": Element{0, "Q_STAR", "", 0, -1, []int8{}},
	"Q_a":    Element{0, "Q_a", "", 0, -1, []int8{}},
	"Q_A":    Element{0, "Q_A", "", 0, -1, []int8{}},
	"Q_AH":   Element{0, "Q_AH", "", 0, -1, []int8{}},
	"Q_M":    Element{0, "Q_M", "", 0, -1, []int8{}},
	"Q_MH":   Element{0, "Q_MH", "", 0, -1, []int8{}},
	"Q_Q":    Element{0, "Q_Q", "", 0, -1, []int8{}},
	"Q_QH":   Element{0, "Q_QH", "", 0, -1, []int8{}},
	"Q_X":    Element{0, "Q_X", "", 0, -1, []int8{}},
	"Q_XH":   Element{0, "Q_XH", "", 0, -1, []int8{}},
	"R":      Element{0, "R", "", 0, -1, []int8{}},
}